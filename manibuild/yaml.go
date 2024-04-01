package manibuild

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func UnmarshalYaml(b []byte, msg proto.Message, m *protojson.UnmarshalOptions) error {
	b, err := yaml.YAMLToJSON(b)
	if err != nil {
		return fmt.Errorf("failed to convert from YAML to JSON: %w", err)
	}

	return m.Unmarshal(b, msg)
}

func MarshalYaml(msg proto.Message, m *protojson.MarshalOptions) ([]byte, error) {
	b, err := m.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return yaml.JSONToYAML(b)
}
