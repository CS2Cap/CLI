package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/output"
)

func newProvidersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "providers",
		Short: "List supported marketplace providers",
		Example: `  cs2cap providers
  cs2cap providers --output json`,
		RunE: func(c *cobra.Command, args []string) error {
			client := newAPIClient()
			resp, err := client.ListProviders(c.Context())
			if err != nil {
				return err
			}

			return renderOutput(output.RowsFunc(func() ([]string, [][]string) {
				if len(resp.Providers) == 0 {
					return nil, nil
				}
				header := []string{"Key", "Name", "Bids", "Sales", "Direct", "Free"}
				rows := make([][]string, 0, len(resp.Providers))
				for _, p := range resp.Providers {
					rows = append(rows, []string{
						p.Key,
						p.Name,
						boolYesNo(p.HasBids),
						boolYesNo(p.HasSales),
						boolYesNo(p.HasDirect),
						boolYesNo(p.IsFree),
					})
				}
				return header, rows
			}))
		},
	}
}

func boolYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func stringInt(n int) string {
	if n == 0 {
		return "0"
	}
	return intString(n)
}

func intString(n int) string {
	if n < 0 {
		return "-" + intString(-n)
	}
	s := ""
	for n >= 1000 {
		s = "," + pad3(n%1000) + s
		n /= 1000
	}
	return itoa(n) + s
}

func pad3(n int) string {
	if n < 10 {
		return "00" + itoa(n)
	}
	if n < 100 {
		return "0" + itoa(n)
	}
	return itoa(n)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	b := make([]byte, 0, 10)
	for n > 0 {
		b = append(b, byte('0'+n%10))
		n /= 10
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
