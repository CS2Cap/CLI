package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/normalize"
	"github.com/cs2cap/cli/internal/output"
)

func newPricesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prices",
		Short: "Query current lowest ask prices",
	}

	listCmd := &cobra.Command{
		Use:   "list [name]",
		Short: "List current lowest ask prices",
		Args:  cobra.MaximumNArgs(1),
		Example: `  cs2cap prices list "AK-47 | Redline FT"
  cs2cap prices list --item-id 1234 --providers steam --providers buff163
  cs2cap prices list "★ Bayonet | Doppler" --phase ruby --currency EUR`,
		RunE: func(c *cobra.Command, args []string) error {
			var params api.ListPricesParams

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
			if phase, _ := c.Flags().GetString("phase"); phase != "" {
				params.Phase = &phase
			}
			params.Providers, _ = c.Flags().GetStringSlice("providers")
			params.Currency, _ = c.Flags().GetString("currency")
			params.Limit, _ = c.Flags().GetInt("limit")
			params.Offset, _ = c.Flags().GetInt("offset")

			client := newAPIClient()
			resp, err := client.ListPrices(c.Context(), params)
			if err != nil {
				return err
			}

			return renderOutput(renderData{
				data: resp,
				toTable: func() ([]string, [][]string) {
					if len(resp.Items) == 0 {
						return nil, nil
					}
					header := []string{"Provider", "Item", "Phase", "Price", "Qty", "Currency"}
					rows := make([][]string, len(resp.Items))
					for i, item := range resp.Items {
						rows[i] = []string{
							item.Provider,
							item.MarketHashName,
							output.Optional(item.Phase),
							output.FormatPrice(item.LowestAsk),
							stringInt(item.Quantity),
							resp.Meta.Currency,
						}
					}
					return header, rows
				},
			})
		},
	}
	listCmd.Flags().Int("item-id", 0, "Filter by item ID")
	listCmd.Flags().String("name", "", "Filter by exact market hash name")
	listCmd.Flags().String("phase", "", "Filter by Doppler phase (ruby, sapphire, etc.)")
	listCmd.Flags().StringSlice("providers", nil, "Filter by provider keys (repeat flag)")
	listCmd.Flags().String("currency", "USD", "Quote currency")
	listCmd.Flags().Int("limit", 20, "Maximum results")
	listCmd.Flags().Int("offset", 0, "Result offset")

	batchCmd := &cobra.Command{
		Use:   "batch",
		Short: "Batch price lookup by item IDs or names",
		Example: `  cs2cap prices batch --items 1,2,3
  cs2cap prices batch --names "AK-47 | Redline FT","★ Bayonet | Doppler"`,
		RunE: func(c *cobra.Command, args []string) error {
			items, _ := c.Flags().GetIntSlice("items")
			rawNames, _ := c.Flags().GetStringSlice("names")
			names := normalize.WearShortcuts(rawNames)

			req := api.BatchParams{
				ItemIDs: items,
				Names:   names,
			}

			client := newAPIClient()
			resp, err := client.BatchPrices(c.Context(), req.ToBatchPrices())
			if err != nil {
				return err
			}

			return renderOutput(renderData{
				data: resp,
				toTable: func() ([]string, [][]string) {
					if len(resp.Items) == 0 {
						return nil, nil
					}
					header := []string{"Item ID", "Name", "Provider", "Price", "Qty"}
					rows := make([][]string, 0)
					for _, item := range resp.Items {
						for _, quote := range item.Quotes {
							rows = append(rows, []string{
								stringInt(item.ItemID),
								item.MarketHashName,
								quote.Provider,
								output.FormatPrice(quote.LowestAsk),
								stringInt(quote.Quantity),
							})
						}
					}
					return header, rows
				},
			})
		},
	}
	batchCmd.Flags().IntSlice("items", nil, "Comma-separated item IDs")
	batchCmd.Flags().StringSlice("names", nil, "Comma-separated market hash names")

	cmd.AddCommand(listCmd)
	cmd.AddCommand(batchCmd)

	return cmd
}
