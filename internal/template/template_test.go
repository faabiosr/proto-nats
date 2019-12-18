package template

import (
	"bytes"
	"testing"
)

func TestTemplateParser(t *testing.T) {
	p := New()

	buf := bytes.NewBuffer(make([]byte, 0))
	data := map[string]interface{}{"Name": ""}

	if err := p.Parse(buf, data); err != nil {
		t.Errorf("parse fails: expected nil - got %v", err)
	}
}

func TestUnexported(t *testing.T) {
	s := "UserGroup"
	expected := "userGroup"

	if unexported(s) != expected {
		t.Errorf("func fail: expected %s - got %s", expected, s)
	}
}
