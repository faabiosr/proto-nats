package cli

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type fakeGen struct{}

func (f *fakeGen) Generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	if len(req.GetFileToGenerate()) == 0 {
		return nil, errors.New("error")
	}

	if len(req.GetFileToGenerate()) == 1 {
		return nil, nil
	}

	return &plugin.CodeGeneratorResponse{}, nil
}

type fakeIO struct{}

func (f *fakeIO) Read(p []byte) (n int, err error) {
	return 0, errors.New("read fail")
}

func (f *fakeIO) Write(p []byte) (n int, err error) {
	return 0, errors.New("write fail")
}

func TestRun(t *testing.T) {
	c := New(&fakeGen{})

	in := bytes.NewReader([]byte("data"))
	out := bytes.NewBuffer(make([]byte, 0))
	fio := &fakeIO{}

	if err := c.Run(fio, out); err == nil {
		t.Errorf("cli fail: expected read error - got %v", err)
	}

	if err := c.Run(in, out); err == nil {
		t.Errorf("cli fail: expected marshal error - got %v", err)
	}

	if err := c.Run(in, out); err == nil {
		t.Errorf("cli fail: expected generate error - got %v", err)
	}

	req := &plugin.CodeGeneratorRequest{FileToGenerate: []string{"1"}}

	data, _ := proto.Marshal(req)
	in = bytes.NewReader(data)

	if err := c.Run(in, out); err == nil {
		t.Errorf("cli fail: expected writer error - got %v", err)
	}

	req.FileToGenerate = append(req.FileToGenerate, "2")

	data, _ = proto.Marshal(req)
	in = bytes.NewReader(data)

	if err := c.Run(in, fio); err == nil {
		t.Errorf("cli fail: expected writer error - got %v", err)
	}

	data, _ = proto.Marshal(req)
	in = bytes.NewReader(data)

	if err := c.Run(in, out); err != nil {
		t.Errorf("cli fail: expected %v - got %v", nil, err)
	}
}
