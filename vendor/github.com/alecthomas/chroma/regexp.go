package chroma

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
)

type Rule struct {
	Pattern string
	Type    Emitter
	Mutator Mutator
}

// An Emitter takes group matches and returns tokens.
type Emitter interface {
	// Emit tokens for the given regex groups.
	Emit(groups []string, lexer Lexer) Iterator
}

// EmitterFunc is a function that is an Emitter.
type EmitterFunc func(groups []string, lexer Lexer) Iterator

// Emit tokens for groups.
func (e EmitterFunc) Emit(groups []string, lexer Lexer) Iterator { return e(groups, lexer) }

// ByGroups emits a token for each matching group in the rule's regex.
func ByGroups(emitters ...Emitter) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer) Iterator {
		iterators := make([]Iterator, 0, len(groups)-1)
		// NOTE: If this panics, there is a mismatch with groups
		for i, group := range groups[1:] {
			iterators = append(iterators, emitters[i].Emit([]string{group}, lexer))
		}
		return Concaterator(iterators...)
	})
}

// UsingByGroup emits tokens for the matched groups in the regex using a
// "sublexer". Used when lexing code blocks where the name of a sublexer is
// contained within the block, for example on a Markdown text block or SQL
// language block.
//
// The sublexer will be retrieved using sublexerGetFunc (typically
// internal.Get), using the captured value from the matched sublexerNameGroup.
//
// If sublexerGetFunc returns a non-nil lexer for the captured sublexerNameGroup,
// then tokens for the matched codeGroup will be emitted using the retrieved
// lexer. Otherwise, if the sublexer is nil, then tokens will be emitted from
// the passed emitter.
//
// Example:
//
//	var Markdown = internal.Register(MustNewLexer(
//		&Config{
//			Name:      "markdown",
//			Aliases:   []string{"md", "mkd"},
//			Filenames: []string{"*.md", "*.mkd", "*.markdown"},
//			MimeTypes: []string{"text/x-markdown"},
//		},
//		Rules{
//			"root": {
//				{"^(```)(\\w+)(\\n)([\\w\\W]*?)(^```$)",
//					UsingByGroup(
//						internal.Get,
//						2, 4,
//						String, String, String, Text, String,
//					),
//					nil,
//				},
//			},
//		},
//	))
//
// See the lexers/m/markdown.go for the complete example.
//
// Note: panic's if the number emitters does not equal the number of matched
// groups in the regex.
func UsingByGroup(sublexerGetFunc func(string) Lexer, sublexerNameGroup, codeGroup int, emitters ...Emitter) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer) Iterator {
		// bounds check
		if len(emitters) != len(groups)-1 {
			panic("UsingByGroup expects number of emitters to be the same as len(groups)-1")
		}

		// grab sublexer
		sublexer := sublexerGetFunc(groups[sublexerNameGroup])

		// build iterators
		iterators := make([]Iterator, len(groups)-1)
		for i, group := range groups[1:] {
			if i == codeGroup-1 && sublexer != nil {
				var err error
				iterators[i], err = sublexer.Tokenise(nil, groups[codeGroup])
				if err != nil {
					panic(err)
				}
			} else {
				iterators[i] = emitters[i].Emit([]string{group}, lexer)
			}
		}

		return Concaterator(iterators...)
	})
}

// Using returns an Emitter that uses a given Lexer for parsing and emitting.
func Using(lexer Lexer) Emitter {
	return EmitterFunc(func(groups []string, _ Lexer) Iterator {
		it, err := lexer.Tokenise(&TokeniseOptions{State: "root", Nested: true}, groups[0])
		if err != nil {
			panic(err)
		}
		return it
	})
}

// UsingSelf is like Using, but uses the current Lexer.
func UsingSelf(state string) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer) Iterator {
		it, err := lexer.Tokenise(&TokeniseOptions{State: state, Nested: true}, groups[0])
		if err != nil {
			panic(err)
		}
		return it
	})
}

// Words creates a regex that matches any of the given literal words.
func Words(prefix, suffix string, words ...string) string {
	for i, word := range words {
		words[i] = regexp.QuoteMeta(word)
	}
	return prefix + `(` + strings.Join(words, `|`) + `)` + suffix
}

// Tokenise text using lexer, returning tokens as a slice.
func Tokenise(lexer Lexer, options *TokeniseOptions, text string) ([]*Token, error) {
	out := []*Token{}
	it, err := lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	for t := it(); t != nil; t = it() {
		out = append(out, t)
	}
	return out, nil
}

// Rules maps from state to a sequence of Rules.
type Rules map[string][]Rule

func (r Rules) Clone() Rules {
	out := map[string][]Rule{}
	for key, rules := range r {
		out[key] = make([]Rule, len(rules))
		copy(out[key], rules)
	}
	return out
}

// MustNewLexer creates a new Lexer or panics.
func MustNewLexer(config *Config, rules Rules) *RegexLexer {
	lexer, err := NewLexer(config, rules)
	if err != nil {
		panic(err)
	}
	return lexer
}

// NewLexer creates a new regex-based Lexer.
//
// "rules" is a state machine transitition map. Each key is a state. Values are sets of rules
// that match input, optionally modify lexer state, and output tokens.
func NewLexer(config *Config, rules Rules) (*RegexLexer, error) {
	if config == nil {
		config = &Config{}
	}
	if _, ok := rules["root"]; !ok {
		return nil, fmt.Errorf("no \"root\" state")
	}
	compiledRules := map[string][]*CompiledRule{}
	for state, rules := range rules {
		compiledRules[state] = nil
		for _, rule := range rules {
			flags := ""
			if !config.NotMultiline {
				flags += "m"
			}
			if config.CaseInsensitive {
				flags += "i"
			}
			if config.DotAll {
				flags += "s"
			}
			compiledRules[state] = append(compiledRules[state], &CompiledRule{Rule: rule, flags: flags})
		}
	}
	return &RegexLexer{
		config: config,
		rules:  compiledRules,
	}, nil
}

func (r *RegexLexer) Trace(trace bool) *RegexLexer {
	r.trace = trace
	return r
}

// A CompiledRule is a Rule with a pre-compiled regex.
//
// Note that regular expressions are lazily compiled on first use of the lexer.
type CompiledRule struct {
	Rule
	Regexp *regexp2.Regexp
	flags  string
}

type CompiledRules map[string][]*CompiledRule

type LexerState struct {
	Lexer *RegexLexer
	Text  []rune
	Pos   int
	Rules CompiledRules
	Stack []string
	State string
	Rule  int
	// Group matches.
	Groups []string
	// Custum context for mutators.
	MutatorContext map[interface{}]interface{}
	iteratorStack  []Iterator
}

func (l *LexerState) Set(key interface{}, value interface{}) {
	l.MutatorContext[key] = value
}

func (l *LexerState) Get(key interface{}) interface{} {
	return l.MutatorContext[key]
}

func (l *LexerState) Iterator() *Token {
	for l.Pos < len(l.Text) && len(l.Stack) > 0 {
		// Exhaust the iterator stack, if any.
		for len(l.iteratorStack) > 0 {
			n := len(l.iteratorStack) - 1
			t := l.iteratorStack[n]()
			if t == nil {
				l.iteratorStack = l.iteratorStack[:n]
				continue
			}
			return t
		}

		l.State = l.Stack[len(l.Stack)-1]
		if l.Lexer.trace {
			fmt.Fprintf(os.Stderr, "%s: pos=%d, text=%q\n", l.State, l.Pos, string(l.Text[l.Pos:]))
		}
		selectedRule, ok := l.Rules[l.State]
		if !ok {
			panic("unknown state " + l.State)
		}
		ruleIndex, rule, groups := matchRules(l.Text[l.Pos:], selectedRule)
		// No match.
		if groups == nil {
			l.Pos++
			return &Token{Error, string(l.Text[l.Pos-1 : l.Pos])}
		}
		l.Rule = ruleIndex
		l.Groups = groups
		l.Pos += utf8.RuneCountInString(groups[0])
		if rule.Mutator != nil {
			if err := rule.Mutator.Mutate(l); err != nil {
				panic(err)
			}
		}
		if rule.Type != nil {
			l.iteratorStack = append(l.iteratorStack, rule.Type.Emit(l.Groups, l.Lexer))
		}
	}
	// Exhaust the IteratorStack, if any.
	// Duplicate code, but eh.
	for len(l.iteratorStack) > 0 {
		n := len(l.iteratorStack) - 1
		t := l.iteratorStack[n]()
		if t == nil {
			l.iteratorStack = l.iteratorStack[:n]
			continue
		}
		return t
	}

	// If we get to here and we still have text, return it as an error.
	if l.Pos != len(l.Text) && len(l.Stack) == 0 {
		value := string(l.Text[l.Pos:])
		l.Pos = len(l.Text)
		return &Token{Type: Error, Value: value}
	}
	return nil
}

type RegexLexer struct {
	config   *Config
	analyser func(text string) float32
	trace    bool

	mu       sync.Mutex
	compiled bool
	rules    map[string][]*CompiledRule
}

// SetAnalyser sets the analyser function used to perform content inspection.
func (r *RegexLexer) SetAnalyser(analyser func(text string) float32) *RegexLexer {
	r.analyser = analyser
	return r
}

func (r *RegexLexer) AnalyseText(text string) float32 {
	if r.analyser != nil {
		return r.analyser(text)
	}
	return 0.0
}

func (r *RegexLexer) Config() *Config {
	return r.config
}

// Regex compilation is deferred until the lexer is used. This is to avoid significant init() time costs.
func (r *RegexLexer) maybeCompile() (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.compiled {
		return nil
	}
	for state, rules := range r.rules {
		for i, rule := range rules {
			if rule.Regexp == nil {
				rule.Regexp, err = regexp2.Compile("^(?"+rule.flags+")(?:"+rule.Pattern+")", 0)
				if err != nil {
					return fmt.Errorf("failed to compile rule %s.%d: %s", state, i, err)
				}
			}
		}
	}
restart:
	seen := map[LexerMutator]bool{}
	for state := range r.rules {
		for i := 0; i < len(r.rules[state]); i++ {
			rule := r.rules[state][i]
			if compile, ok := rule.Mutator.(LexerMutator); ok {
				if seen[compile] {
					return fmt.Errorf("saw mutator %T twice; this should not happen", compile)
				}
				seen[compile] = true
				if err := compile.MutateLexer(r.rules, state, i); err != nil {
					return err
				}
				// Process the rules again in case the mutator added/removed rules.
				//
				// This sounds bad, but shouldn't be significant in practice.
				goto restart
			}
		}
	}
	r.compiled = true
	return nil
}

func (r *RegexLexer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	if err := r.maybeCompile(); err != nil {
		return nil, err
	}
	if options == nil {
		options = defaultOptions
	}
	if !options.Nested && r.config.EnsureNL && !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	state := &LexerState{
		Lexer:          r,
		Text:           []rune(text),
		Stack:          []string{options.State},
		Rules:          r.rules,
		MutatorContext: map[interface{}]interface{}{},
	}
	return state.Iterator, nil
}

func matchRules(text []rune, rules []*CompiledRule) (int, *CompiledRule, []string) {
	for i, rule := range rules {
		match, err := rule.Regexp.FindRunesMatch(text)
		if match != nil && err == nil {
			groups := []string{}
			for _, g := range match.Groups() {
				groups = append(groups, g.String())
			}
			return i, rule, groups
		}
	}
	return 0, &CompiledRule{}, nil
}
