package template

import (
	"io"
	"io/ioutil"
	tpl "text/template"

	"github.com/faabiosr/proto-nats/parser"
	"github.com/markbates/pkger"
)

type template struct{}

// New creates an instance of Template Parser.
func New() parser.Parser {
	return &template{}
}

// Parser parses the template file provided by pkger into io.Writer.
func (t *template) Parse(w io.Writer, data interface{}) error {
	f, err := pkger.Open("/templates/resource.tmpl")

	if err != nil {
		return err
	}

	defer f.Close()

	content, err := ioutil.ReadAll(f)

	if err != nil {
		return nil
	}

	tpl, err := tpl.New("proto-nats").Parse(string(content))

	if err != nil {
		return err
	}

	return tpl.Execute(w, data)
}