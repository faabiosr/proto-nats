package template

import (
	"testing"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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

func TestSubject(t *testing.T) {
	opts := &descriptor.MethodOptions{}
	v := subject(opts)

	if v != nil {
		t.Error("func fail: expected a nil value")
	}
}

func TestSubjectPrefix(t *testing.T) {
	opts := &descriptor.ServiceOptions{}
	v := subjectPrefix(opts)

	if v != nil {
		t.Error("func fail: expected a nil value")
	}
}
