package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/normalize"
	"github.com/cs2cap/cli/internal/output"
)

func newSalesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sales",
		Short: "Query recent sales history",
	}

	listCmd := &cobra.Command{
		Use:   "list [name]",
		Short: "List recent sales for an item",
		Args:  cobra.MaximumNArgs(1),
		Example: `  cs2cap sales list "AK-47 | Redline FT"
  cs2cap sales list --item-id 1234 --providers steam --limit 10`,
		RunE: func(c *cobra.Command, args []string) error {
			var params api.ListSalesParams

			if itemID, _ := c.Flags().GetInt("item-id"); c.Flags().Changed("item-id") {
				params.ItemID = &itemID
			}
			flagName, _ := c.Flags().GetString("name")
			name := flagName
			if name == "" && len(args) > 0 {
				name = args[0]
			}
			if name != "" {
				expanded := normalize.WearShortcut(name)
				params.MarketHashName = &expanded
			}
			params.Providers, _ = c.Flags().GetStringSlice("providers")
			params.Limit, _ = c.Flags().GetInt("limit")

			client := newAPIClient()
			resp, err := client.ListSales(c.Context(), params)
			if err != nil {
				return err
			}

			return renderOutput(renderData{
				data: resp,
				toTable: func() ([]string, [][]string) {
					if len(resp.Items) == 0 {
						return nil, nil
					}
					header := []string{"Provider", "Date", "Item", "Price", "Float"}
					rows := make([][]string, len(resp.Items))
					for i, sale := range resp.Items {
						rows[i] = []string{
							sale.Provider,
							sale.Date,
							sale.MarketHashName,
							output.FormatPrice(sale.Price),
							output.OptionalFloat(sale.Float),
						}
					}
					return header, rows
				},
			})
		},
	}
	listCmd.Flags().Int("item-id", 0, "Filter by item ID")
	listCmd.Flags().String("name", "", "Filter by exact market hash name")
	listCmd.Flags().StringSlice("providers", nil, "Filter by provider keys (repeat flag)")
	listCmd.Flags().Int("limit", 20, "Maximum results")

	cmd.AddCommand(listCmd)

	return cmd
}
