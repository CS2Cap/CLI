package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/normalize"
	"github.com/cs2cap/cli/internal/output"
)

func newBidsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bids",
		Short: "Query current highest buy orders",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List current highest buy orders",
		Example: `  cs2cap bids list --name "AK-47 | Redline (Field-Tested)"
  cs2cap bids list --name "AK-47 | Redline FT"          # wear shortcut
  cs2cap bids list --item-id 1234 --providers steam --providers buff163`,
		RunE: func(c *cobra.Command, args []string) error {
			var params api.ListBidsParams

			if itemID, _ := c.Flags().GetInt("item-id"); c.Flags().Changed("item-id") {
				params.ItemID = &itemID
			}
			if name, _ := c.Flags().GetString("name"); name != "" {
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
			resp, err := client.ListBids(c.Context(), params)
			if err != nil {
				return err
			}

			return renderOutput(output.RowsFunc(func() ([]string, [][]string) {
				if len(resp.Items) == 0 {
					return nil, nil
				}
				header := []string{"Provider", "Item", "Phase", "Bid", "Orders", "Currency"}
				rows := make([][]string, len(resp.Items))
				for i, item := range resp.Items {
					rows[i] = []string{
						item.Provider,
						item.MarketHashName,
						output.Optional(item.Phase),
						output.FormatPrice(item.HighestBid),
						stringInt(item.NumBids),
						resp.Meta.Currency,
					}
				}
				return header, rows
			}))
		},
	}
	listCmd.Flags().Int("item-id", 0, "Filter by item ID")
	listCmd.Flags().String("name", "", "Filter by exact market hash name")
	listCmd.Flags().String("phase", "", "Filter by Doppler phase")
	listCmd.Flags().StringSlice("providers", nil, "Filter by provider keys (repeat flag)")
	listCmd.Flags().String("currency", "USD", "Quote currency")
	listCmd.Flags().Int("limit", 20, "Maximum results")
	listCmd.Flags().Int("offset", 0, "Result offset")

	batchCmd := &cobra.Command{
		Use:   "batch",
		Short: "Batch bid lookup by item IDs or names",
		Example: `  cs2cap bids batch --items 1,2,3
  cs2cap bids batch --names "AK-47 | Redline FT","★ Bayonet | Doppler"`,
		RunE: func(c *cobra.Command, args []string) error {
			items, _ := c.Flags().GetIntSlice("items")
			rawNames, _ := c.Flags().GetStringSlice("names")
			names := normalize.WearShortcuts(rawNames)

			client := newAPIClient()
			resp, err := client.BatchBids(c.Context(), api.BatchParams{ItemIDs: items, Names: names}.ToBatchBids())
			if err != nil {
				return err
			}

			return renderOutput(output.RowsFunc(func() ([]string, [][]string) {
				if len(resp.Items) == 0 {
					return nil, nil
				}
				header := []string{"Item ID", "Name", "Provider", "Bid", "Orders"}
				rows := make([][]string, 0)
				for _, item := range resp.Items {
					for _, quote := range item.Quotes {
						rows = append(rows, []string{
							stringInt(item.ItemID),
							item.MarketHashName,
							quote.Provider,
							output.FormatPrice(quote.HighestBid),
							stringInt(quote.NumBids),
						})
					}
				}
				return header, rows
			}))
		},
	}
	batchCmd.Flags().IntSlice("items", nil, "Comma-separated item IDs")
	batchCmd.Flags().StringSlice("names", nil, "Comma-separated market hash names")

	cmd.AddCommand(listCmd)
	cmd.AddCommand(batchCmd)

	return cmd
}
