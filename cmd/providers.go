package cmd

import (
	"sort"

	"github.com/spf13/cobra"
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

			return renderOutput(renderData{
				data: resp,
				toTable: func() ([]string, [][]string) {
					if len(resp) == 0 {
						return nil, nil
					}

					type entry struct {
						name string
						info struct {
							Key   string
							Code  string
							Type  string
							Curr  string
							Bids  bool
							Sales bool
							Up    bool
						}
					}

					keys := make([]string, 0, len(resp))
					for name := range resp {
						keys = append(keys, name)
					}
					sort.Strings(keys)

					header := []string{"Name", "Key", "Code", "Type", "Currency", "Bids", "Sales", "Status"}
					rows := make([][]string, 0, len(resp))
					for _, name := range keys {
						p := resp[name]
						status := p.Health.Status
						if status == "" {
							status = "unknown"
						}
						rows = append(rows, []string{
							name,
							p.Key,
							p.Code,
							p.MarketType,
							p.DefaultCurrency,
							boolYesNo(p.Features.HasBuyOrders),
							boolYesNo(p.Features.HasRecentSales),
							status,
						})
					}
					return header, rows
				},
			})
		},
	}
}
