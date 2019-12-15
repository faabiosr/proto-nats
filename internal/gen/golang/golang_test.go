package golang

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func TestName(t *testing.T) {
	if name := Name("proto/user.proto"); name != "user" {
		t.Errorf("name fail: expected %s - got %s", "user", name)
	}
}

func TestPackage(t *testing.T) {
	fd := &descriptor.FileDescriptorProto{
		Package: proto.String("pkg_user"),
	}

	if pkg := Package(fd); pkg != "pkg_user" {
		t.Errorf("package fail: expected %s - got %s", "pkg_user", pkg)
	}

	fd.Options = &descriptor.FileOptions{
		GoPackage: proto.String("user"),
	}

	if pkg := Package(fd); pkg != "user" {
		t.Errorf("package fail: expected %s - got %s", "user", pkg)
	}
}
