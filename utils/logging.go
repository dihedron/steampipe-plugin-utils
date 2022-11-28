package utils

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ToJSON dumps the input object to JSON.
func ToJSON(v any) string {
	s, _ := json.Marshal(v)
	return string(s)
}

// ToPrettyJSON dumps the input object to JSON.
func ToPrettyJSON(v any) string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

// ToYAML dumps the input object to YAML.
func ToYAML(v any) string {
	s, _ := yaml.Marshal(v)
	return string(s)
}

// PointerTo returns a pointer to a given value.
func PointerTo[T any](value T) *T {
	return &value
}
