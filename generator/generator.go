package generator

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Generator represents an interface for generating code.
type Generator interface {
	// Generate generates the code package.
	Generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error)
}
