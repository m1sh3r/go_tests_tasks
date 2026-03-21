package golden

import (
	"sort"
	"strconv"
	"strings"
)

type CSVResult struct {
	ID     string
	A      float64
	B      float64
	C      float64
	Root1  *float64
	Root2  *float64
	Status string
}

func ResultsToCSV(results []CSVResult) string {
	cp := make([]CSVResult, len(results))
	copy(cp, results)
	sort.Slice(cp, func(i, j int) bool { return cp[i].ID < cp[j].ID })

	var b strings.Builder
	b.WriteString("id,a,b,c,root1,root2,status\n")
	for _, r := range cp {
		row := []string{
			escapeCSV(r.ID),
			escapeCSV(trimFloat(r.A)),
			escapeCSV(trimFloat(r.B)),
			escapeCSV(trimFloat(r.C)),
			escapeCSV(formatOptFloat(r.Root1)),
			escapeCSV(formatOptFloat(r.Root2)),
			escapeCSV(r.Status),
		}
		b.WriteString(strings.Join(row, ","))
		b.WriteString("\n")
	}
	return strings.TrimRight(b.String(), "\r\n")
}

func formatOptFloat(v *float64) string {
	if v == nil {
		return ""
	}
	return trimFloat(*v)
}

func escapeCSV(v string) string {
	if strings.ContainsAny(v, "\",\n\r") {
		return `"` + strings.ReplaceAll(v, `"`, `""`) + `"`
	}
	return v
}

func makePtr(v float64) *float64 {
	return &v
}

func formatInt(v int) string {
	return strconv.Itoa(v)
}
