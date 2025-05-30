package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/invopop/jsonschema"
)

type ToolDefinition struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	InputSchema anthropic.ToolInputSchemaParam `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}

func GenerateSchema[T any]() anthropic.ToolInputSchemaParam {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T

	schema := reflector.Reflect(v)

	return anthropic.ToolInputSchemaParam{
		Properties: schema.Properties,
	}
}

// ValidatePath returns an error if `path` is not allowed to be read
func ValidatePath(path string) error {
	if IsFileIgnored(filepath.Base(path)) {
		return fmt.Errorf("permission denied for file: %s", path)
	}

	illegalPrefixes := []string{"/", ".."}
	for _, prefix := range illegalPrefixes {
		if strings.HasPrefix(path, prefix) {
			return fmt.Errorf("permission denied for prefix: %s", prefix)
		}
	}
	return nil
}
