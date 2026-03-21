package fuzzing

import (
	"strconv"
	"testing"
)

func FuzzParseFloatSafe(f *testing.F) {
	f.Add("")
	f.Add(" ")
	f.Add("NaN")
	f.Add("Inf")
	f.Add("-Inf")
	f.Add("3.14")
	f.Add("1e309")
	f.Add("abc")

	f.Fuzz(func(t *testing.T, s string) {
		_, _ = strconv.ParseFloat(s, 64)
		_, _ = ParseFloatSafe(s)
	})
}
