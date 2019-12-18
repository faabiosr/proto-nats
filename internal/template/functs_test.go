package template

import (
	"testing"
)

func TestUnexported(t *testing.T) {
	s := "UserGroup"
	expected := "userGroup"

	if unexported(s) != expected {
		t.Errorf("func fail: expected %s - got %s", expected, s)
	}
}

func TestDeref(t *testing.T) {
	expected := "deref"
	s := deref(&expected)

	if expected != s.(string) {
		t.Errorf("func fail: expected a deref value")
	}

	s = deref(expected)

	if expected != s.(string) {
		t.Errorf("func fail: expected a deref value")
	}
}
