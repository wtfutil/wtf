# Chroma - A general purpose syntax highlighter in pure Go [![](https://godoc.org/github.com/alecthomas/chroma?status.svg)](http://godoc.org/github.com/alecthomas/chroma) [![Build Status](https://travis-ci.org/alecthomas/chroma.png)](https://travis-ci.org/alecthomas/chroma) [![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

> **NOTE:** As Chroma has just been released, its API is still in flux. That said, the high-level interface should not change significantly.

Chroma takes source code and other structured text and converts it into syntax
highlighted HTML, ANSI-coloured text, etc.

Chroma is based heavily on [Pygments](http://pygments.org/), and includes
translators for Pygments lexers and styles.

## Table of Contents

<!-- MarkdownTOC -->

1. [Supported languages](#supported-languages)
1. [Using the library](#using-the-library)
    1. [Quick start](#quick-start)
    1. [Identifying the language](#identifying-the-language)
    1. [Formatting the output](#formatting-the-output)
    1. [The HTML formatter](#the-html-formatter)
1. [More detail](#more-detail)
    1. [Lexers](#lexers)
    1. [Formatters](#formatters)
    1. [Styles](#styles)
1. [Command-line interface](#command-line-interface)
1. [What's missing compared to Pygments?](#whats-missing-compared-to-pygments)

<!-- /MarkdownTOC -->

## Supported languages

ABNF, ANTLR, APL, ActionScript, ActionScript 3, Ada, Angular2, ApacheConf, AppleScript, Awk, BNF, Ballerina, Base Makefile, Bash, Batchfile, BlitzBasic, Brainfuck, C, C#, C++, CFEngine3, CMake, COBOL, CSS, Cap'n Proto, Ceylon, ChaiScript, Cheetah, Clojure, CoffeeScript, Common Lisp, Coq, Crystal, Cython, DTD, Dart, Diff, Django/Jinja, Docker, EBNF, Elixir, Elm, EmacsLisp, Erlang, FSharp, Factor, Fish, Forth, Fortran, GAS, GDScript, GLSL, Genshi, Genshi HTML, Genshi Text, Gnuplot, Go, Go HTML Template, Go Text Template, Groovy, HTML, HTTP, Handlebars, Haskell, Haxe, Hexdump, Hy, INI, Idris, Io, JSON, JSX, Java, JavaScript, Julia, Kotlin, LLVM, Lighttpd configuration file, Lua, Mako, Mason, Mathematica, MiniZinc, Modula-2, MorrowindScript, MySQL, Myghty, NASM, Newspeak, Nginx configuration file, Nim, Nix, OCaml, Objective-C, Octave, Org Mode, PHP, PL/pgSQL, POVRay, PacmanConf, Perl, Pig, PkgConfig, PostScript, PostgreSQL SQL dialect, PowerShell, Prolog, Protocol Buffer, Puppet, Python, Python 3, QBasic, R, Racket, Ragel, Rexx, Ruby, Rust, SCSS, SPARQL, SQL, Sass, Scala, Scheme, Scilab, Smalltalk, Smarty, Snobol, Solidity, SquidConf, Swift, TASM, TOML, Tcl, Tcsh, TeX, Termcap, Terminfo, Terraform, Thrift, Transact-SQL, Turtle, Twig, TypeScript, TypoScript, TypoScriptCssData, TypoScriptHtmlData, VHDL, VimL, WDTE, XML, Xorg, YAML, cfstatement, markdown, reStructuredText, reg, systemverilog, verilog

_I will attempt to keep this section up to date, but an authoritative list can be
displayed with `chroma --list`._

## Using the library

Chroma, like Pygments, has the concepts of
[lexers](https://github.com/alecthomas/chroma/tree/master/lexers),
[formatters](https://github.com/alecthomas/chroma/tree/master/formatters) and
[styles](https://github.com/alecthomas/chroma/tree/master/styles).

Lexers convert source text into a stream of tokens, styles specify how token
types are mapped to colours, and formatters convert tokens and styles into
formatted output.

A package exists for each of these, containing a global `Registry` variable
with all of the registered implementations. There are also helper functions
for using the registry in each package, such as looking up lexers by name or
matching filenames, etc.

In all cases, if a lexer, formatter or style can not be determined, `nil` will
be returned. In this situation you may want to default to the `Fallback`
value in each respective package, which provides sane defaults.

### Quick start

A convenience function exists that can be used to simply format some source
text, without any effort:

```go
err := quick.Highlight(os.Stdout, someSourceCode, "go", "html", "monokai")
```

### Identifying the language

To highlight code, you'll first have to identify what language the code is
written in. There are three primary ways to do that:

1. Detect the language from its filename.

    ```go
    lexer := lexers.Match("foo.go")
    ```

3. Explicitly specify the language by its Chroma syntax ID (a full list is available from `lexers.Names()`).

    ```go
    lexer := lexers.Get("go")
    ```

3. Detect the language from its content.

    ```go
    lexer := lexers.Analyse("package main\n\nfunc main()\n{\n}\n")
    ```

In all cases, `nil` will be returned if the language can not be identified.

```go
if lexer == nil {
  lexer = lexers.Fallback
}
```

At this point, it should be noted that some lexers can be extremely chatty. To
mitigate this, you can use the coalescing lexer to coalesce runs of identical
token types into a single token:

```go
lexer = chroma.Coalesce(lexer)
```

### Formatting the output

Once a language is identified you will need to pick a formatter and a style (theme).

```go
style := styles.Get("swapoff")
if style == nil {
  style = styles.Fallback
}
formatter := formatters.Get("html")
if formatter == nil {
  formatter = formatters.Fallback
}
```

Then obtain an iterator over the tokens:

```go
contents, err := ioutil.ReadAll(r)
iterator, err := lexer.Tokenise(nil, string(contents))
```

And finally, format the tokens from the iterator:

```go
err := formatter.Format(w, style, iterator)
```

### The HTML formatter

By default the `html` registered formatter generates standalone HTML with
embedded CSS. More flexibility is available through the `formatters/html` package.

Firstly, the output generated by the formatter can be customised with the
following constructor options:

- `Standalone()` - generate standalone HTML with embedded CSS.
- `WithClasses()` - use classes rather than inlined style attributes.
- `ClassPrefix(prefix)` - prefix each generated CSS class.
- `TabWidth(width)` - Set the rendered tab width, in characters.
- `WithLineNumbers()` - Render line numbers (style with `LineNumbers`).
- `HighlightLines(ranges)` - Highlight lines in these ranges (style with `LineHighlight`).
- `LineNumbersInTable()` - Use a table for formatting line numbers and code, rather than spans.

If `WithClasses()` is used, the corresponding CSS can be obtained from the formatter with:

```go
formatter := html.New(html.WithClasses())
err := formatter.WriteCSS(w, style)
```

## More detail

### Lexers

See the [Pygments documentation](http://pygments.org/docs/lexerdevelopment/)
for details on implementing lexers. Most concepts apply directly to Chroma,
but see existing lexer implementations for real examples.

In many cases lexers can be automatically converted directly from Pygments by
using the included Python 3 script `pygments2chroma.py`. I use something like
the following:

```
python3 ~/Projects/chroma/_tools/pygments2chroma.py \
  pygments.lexers.jvm.KotlinLexer \
  > ~/Projects/chroma/lexers/kotlin.go \
  && gofmt -s -w ~/Projects/chroma/lexers/*.go
```

See notes in [pygments-lexers.go](https://github.com/alecthomas/chroma/blob/master/pygments-lexers.txt)
for a list of lexers, and notes on some of the issues importing them.

### Formatters

Chroma supports HTML output, as well as terminal output in 8 colour, 256 colour, and true-colour.

A `noop` formatter is included that outputs the token text only, and a `tokens`
formatter outputs raw tokens. The latter is useful for debugging lexers.

### Styles

Chroma styles use the [same syntax](http://pygments.org/docs/styles/) as Pygments.

All Pygments styles have been converted to Chroma using the `_tools/style.py` script.

For a quick overview of the available styles and how they look, check out the [Chroma Style Gallery](https://xyproto.github.io/splash/docs/).

## Command-line interface

A command-line interface to Chroma is included. It can be installed with:

```
go get -u github.com/alecthomas/chroma/cmd/chroma
```

## What's missing compared to Pygments?

- Quite a few lexers, for various reasons (pull-requests welcome):
    - Pygments lexers for complex languages often include custom code to
      handle certain aspects, such as Perl6's ability to nest code inside
      regular expressions. These require time and effort to convert.
    - I mostly only converted languages I had heard of, to reduce the porting cost.
- Some more esoteric features of Pygments are omitted for simplicity.
- Though the Chroma API supports content detection, very few languages support them.
  I have plans to implement a statistical analyser at some point, but not enough time.
