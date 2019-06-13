package armory

import (
	"bytes"
	"encoding/json"
)

type j struct{}

// String string
var Json *j

// MarshalWithoutEscapeHTML MarshalWithoutEscapeHTML
func (s *j) MarshalWithoutEscapeHTML(v interface{}, pretty bool) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if pretty {
		encoder.SetIndent("", "    ")
	}
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return bytes.TrimSpace(buffer.Bytes()), err
}

// Indent Indent
func (s *j) Indent(bts []byte) ([]byte, error) {
	out := bytes.Buffer{}
	err := json.Indent(&out, bts, "", "  ")
	return out.Bytes(), err
}
