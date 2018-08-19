package chroma

// An Iterator across tokens.
//
// nil will be returned at the end of the Token stream.
//
// If an error occurs within an Iterator, it may propagate this in a panic. Formatters should recover.
type Iterator func() *Token

// Tokens consumes all tokens from the iterator and returns them as a slice.
func (i Iterator) Tokens() []*Token {
	out := []*Token{}
	for t := i(); t != nil; t = i() {
		out = append(out, t)
	}
	return out
}

// Concaterator concatenates tokens from a series of iterators.
func Concaterator(iterators ...Iterator) Iterator {
	return func() *Token {
		for len(iterators) > 0 {
			t := iterators[0]()
			if t != nil {
				return t
			}
			iterators = iterators[1:]
		}
		return nil
	}
}

// Literator converts a sequence of literal Tokens into an Iterator.
func Literator(tokens ...*Token) Iterator {
	return func() (out *Token) {
		if len(tokens) == 0 {
			return nil
		}
		token := tokens[0]
		tokens = tokens[1:]
		return token
	}
}
