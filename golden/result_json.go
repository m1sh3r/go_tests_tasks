package golden

import (
	"bytes"
	"encoding/json"
)

type Result struct {
	ID     string    `json:"id"`
	A      float64   `json:"a"`
	B      float64   `json:"b"`
	C      float64   `json:"c"`
	Roots  []float64 `json:"roots"`
	Status string    `json:"status"`
}

func ResultToJSON(r Result) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	_ = enc.Encode(r)
	return bytes.TrimRight(buf.Bytes(), "\r\n")
}
