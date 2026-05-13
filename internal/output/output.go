package output

import (
	"fmt"
	"io"
	"strings"
)

type Format int

const (
	FormatTable Format = iota
	FormatJSON
)

func ParseFormat(s string) (Format, error) {
	switch strings.ToLower(s) {
	case "table":
		return FormatTable, nil
	case "json":
		return FormatJSON, nil
	default:
		return FormatTable, fmt.Errorf("unknown output format: %q (supported: table, json)", s)
	}
}

func FormatPrice(cents int) string {
	if cents < 0 {
		return fmt.Sprintf("-$%d.%02d", -cents/100, -cents%100)
	}
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}

func Optional(s *string) string {
	if s == nil {
		return "-"
	}
	return *s
}

func OptionalInt(i *int) string {
	if i == nil {
		return "-"
	}
	return fmt.Sprintf("%d", *i)
}

func OptionalFloat(f *float64) string {
	if f == nil {
		return "-"
	}
	return fmt.Sprintf("%.4f", *f)
}

func OptionalBool(b *bool) string {
	if b == nil {
		return "-"
	}
	if *b {
		return "yes"
	}
	return "no"
}

type Renderer interface {
	Render(w io.Writer, v interface{}) error
}
