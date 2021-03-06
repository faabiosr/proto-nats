package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"

	"github.com/faabiosr/proto-nats/generator"
	"github.com/faabiosr/proto-nats/internal/gen/golang"
	"github.com/faabiosr/proto-nats/parser"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type gen struct {
	version string
	parser  parser.Parser
}

// ErrNoFiles returns an error when wasn't files to generate.
var ErrNoFiles = errors.New("no files to generate")

// New creates an instance of generator.
func New(version string, parser parser.Parser) generator.Generator {
	return &gen{version, parser}
}

// Generator generates the Golang code.
func (g *gen) Generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	if len(req.GetFileToGenerate()) == 0 {
		return nil, ErrNoFiles
	}

	files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.GetProtoFile()))

	for _, fd := range req.GetProtoFile() {
		if len(fd.Service) == 0 {
			continue
		}

		file, err := g.file(fd)

		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return &plugin.CodeGeneratorResponse{File: files}, nil
}

// file builds the golang code based on parser.
func (g *gen) file(fd *descriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	name := fmt.Sprintf("%s.%s.go", golang.Name(fd.GetName()), "nats")

	data := map[string]interface{}{
		"File":     name,
		"Package":  golang.Package(fd),
		"Services": fd.Service,
		"Version":  g.version,
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	if err := g.parser.Parse(buf, data); err != nil {
		return nil, err
	}

	content, err := format.Source(buf.Bytes())

	if err != nil {
		return nil, err
	}

	res := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(name),
		Content: proto.String(string(content)),
	}

	return res, nil
}
