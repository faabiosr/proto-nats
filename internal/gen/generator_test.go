package gen

import (
	"io"
	"testing"

	"github.com/faabiosr/proto-nats/parser"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type fakeParser struct {
}

func (fp *fakeParser) Parse(out io.Writer, data interface{}) error {
	d := data.(map[string]interface{})
	pkg := d["Package"].(string)

	if pkg == "user" {
		return parser.ErrParseFail
	}

	return nil
}

func TestGenerate(t *testing.T) {
	g := New("v0.0.0", &fakeParser{})

	req := &plugin.CodeGeneratorRequest{}

	if _, err := g.Generate(req); err != ErrNoFiles {
		t.Errorf("generator fail: expected %v - got %v", ErrNoFiles, err)
	}

	req.FileToGenerate = []string{"user.proto"}

	res, _ := g.Generate(req)

	if total := len(res.GetFile()); total > 0 {
		t.Errorf("generator fail: should generate %d instead of %d", 0, total)
	}

	req.ProtoFile = []*descriptor.FileDescriptorProto{
		&descriptor.FileDescriptorProto{
			Name:    proto.String("user"),
			Package: proto.String("user"),
		},
	}

	if _, err := g.Generate(req); err == nil {
		t.Errorf("generate fail: expected an error - got %v", err)
	}

	req.ProtoFile = []*descriptor.FileDescriptorProto{
		&descriptor.FileDescriptorProto{
			Name:    proto.String("group"),
			Package: proto.String("group"),
		},
	}

	res, _ = g.Generate(req)

	if total := len(res.GetFile()); total == 0 {
		t.Errorf("generator fail: should generate %d instead of %d", total, 0)
	}
}
