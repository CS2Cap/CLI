package output

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type TableRenderer struct{}

func NewTableRenderer() *TableRenderer {
	return &TableRenderer{}
}

type RowsFunc func() (header []string, rows [][]string)

func (t *TableRenderer) Render(w io.Writer, v interface{}) error {
	fn, ok := v.(RowsFunc)
	if !ok {
		return fmt.Errorf("table renderer expects a RowsFunc, got %T", v)
	}

	header, rows := fn()
	if header == nil {
		_, err := fmt.Fprintln(w, "No results.")
		return err
	}

	tw := tabwriter.NewWriter(w, 0, 0, 3, ' ', 0)

	for i, h := range header {
		if i > 0 {
			fmt.Fprint(tw, "\t")
		}
		fmt.Fprint(tw, h)
	}
	fmt.Fprintln(tw)

	seps := make([]string, len(header))
	for i, h := range header {
		seps[i] = strings.Repeat("-", len(h))
	}
	fmt.Fprintln(tw, strings.Join(seps, "\t"))

	for _, row := range rows {
		fmt.Fprintln(tw, strings.Join(row, "\t"))
	}

	return tw.Flush()
}
