package cli

import (
	"io"
	"io/ioutil"

	"github.com/faabiosr/proto-nats/generator"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type cli struct {
	gen generator.Generator
}

// New creates a new instance of CLI.
func New(gen generator.Generator) *cli {
	return &cli{gen}
}

// Run runs the cli reading and writing the generated code.
func (c *cli) Run(in io.Reader, out io.Writer) error {
	req, err := c.read(in)

	if err != nil {
		return err
	}

	res, err := c.gen.Generate(req)

	if err != nil {
		return err
	}

	return c.write(out, res)
}

func (c *cli) read(in io.Reader) (*plugin.CodeGeneratorRequest, error) {
	data, err := ioutil.ReadAll(in)

	if err != nil {
		return nil, err
	}

	req := &plugin.CodeGeneratorRequest{}

	if err := proto.Unmarshal(data, req); err != nil {
		return nil, err
	}

	return req, nil
}

func (c *cli) write(out io.Writer, res *plugin.CodeGeneratorResponse) error {
	data, err := proto.Marshal(res)

	if err != nil {
		return err
	}

	if _, err := out.Write(data); err != nil {
		return err
	}

	return nil
}
