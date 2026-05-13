package output

import (
	"encoding/json"
	"io"
)

type JSONRenderer struct{}

func NewJSONRenderer() *JSONRenderer {
	return &JSONRenderer{}
}

func (r *JSONRenderer) Render(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
