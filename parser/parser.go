package parser

import "io"

// Parser represents a template parser.
type Parser interface {
	// Parser parses the template file into io.Writer.
	Parse(w io.Writer, data interface{}) error
}
