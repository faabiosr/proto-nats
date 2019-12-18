package template

import (
	"reflect"
	"strings"

	pn "github.com/faabiosr/proto-nats/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func unexported(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func deref(v interface{}) interface{} {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr {
		return v
	}

	return val.Elem().Interface()
}

func subject(opts *descriptor.MethodOptions) interface{} {
	opt, _ := proto.GetExtension(opts, pn.E_Subject)

	return opt
}

func subjectPrefix(opts *descriptor.ServiceOptions) interface{} {
	opt, _ := proto.GetExtension(opts, pn.E_SubjectPrefix)

	return opt
}
