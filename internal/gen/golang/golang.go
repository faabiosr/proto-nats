package golang

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Name returns the name of Golang file.
func Name(s string) string {
	if i := strings.LastIndex(s, "/"); i >= 0 {
		s = s[i+1:]
	}

	if i := strings.LastIndex(s, "."); i >= 0 {
		s = s[0:i]
	}

	return s
}

// Package returns the package name of Golang file.
func Package(fd *descriptor.FileDescriptorProto) string {
	if opt := fd.GetOptions(); opt != nil {
		if pkg := opt.GetGoPackage(); pkg != "" {
			return Name(pkg)
		}
	}

	return Name(fd.GetPackage())
}
