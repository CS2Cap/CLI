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
		Use:   "list",
		Short: "List recent sales for an item",
		Example: `  cs2cap sales list --name "AK-47 | Redline (Field-Tested)"
  cs2cap sales list --name "AK-47 | Redline FT"         # wear shortcut
  cs2cap sales list --item-id 1234 --providers steam --limit 10`,
		RunE: func(c *cobra.Command, args []string) error {
			var params api.ListSalesParams

			if itemID, _ := c.Flags().GetInt("item-id"); c.Flags().Changed("item-id") {
				params.ItemID = &itemID
			}
			if name, _ := c.Flags().GetString("name"); name != "" {
				expanded := normalize.WearShortcut(name)
				params.MarketHashName = &expanded
			}
			params.Providers, _ = c.Flags().GetStringSlice("providers")
			params.Limit, _ = c.Flags().GetInt("limit")
			params.Offset, _ = c.Flags().GetInt("offset")

			client := newAPIClient()
			resp, err := client.ListSales(c.Context(), params)
			if err != nil {
				return err
			}

			return renderOutput(output.RowsFunc(func() ([]string, [][]string) {
				if len(resp.Items) == 0 {
					return nil, nil
				}
				header := []string{"Provider", "Sale ID", "Item", "Price", "Sold At"}
				rows := make([][]string, 0)
				for _, prov := range resp.Items {
					for _, sale := range prov.Sales {
						rows = append(rows, []string{
							prov.Provider,
							sale.SaleID,
							sale.MarketHashName,
							output.FormatPrice(sale.Price),
							sale.SoldAt,
						})
					}
				}
				return header, rows
			}))
		},
	}
	listCmd.Flags().Int("item-id", 0, "Filter by item ID")
	listCmd.Flags().String("name", "", "Filter by exact market hash name")
	listCmd.Flags().StringSlice("providers", nil, "Filter by provider keys (repeat flag)")
	listCmd.Flags().Int("limit", 20, "Maximum results")
	listCmd.Flags().Int("offset", 0, "Result offset")

	cmd.AddCommand(listCmd)

	return cmd
}
