package parser

import (
	"errors"
	"io"
)

var ErrParseFail = errors.New("parse fails")

// Parser represents a template parser.
type Parser interface {
	// Parser parses the template file into io.Writer.
	Parse(w io.Writer, data interface{}) error
}
