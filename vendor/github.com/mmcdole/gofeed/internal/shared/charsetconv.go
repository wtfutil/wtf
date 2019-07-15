package shared

import (
	"io"

	"golang.org/x/net/html/charset"
)

func NewReaderLabel(label string, input io.Reader) (io.Reader, error) {
	conv, err := charset.NewReaderLabel(label, input)

	if err != nil {
		return nil, err
	}

	// Wrap the charset decoder reader with a XML sanitizer
	//clean := NewXMLSanitizerReader(conv)
	return conv, nil
}
