package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/output"
)

func newItemsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "items",
		Short: "Query the CS2 item catalog",
	}

	searchCmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Search or filter the item catalog",
		Args:  cobra.MaximumNArgs(1),
		Example: `  cs2cap items search "AK-47"
  cs2cap items search --type weapon --rarity "Covert"
  cs2cap items search --type sticker --limit 50`,
		RunE: func(c *cobra.Command, args []string) error {
			flagQ := c.Flags().Lookup("q").Value.String()
			q := flagQ
			if q == "" && len(args) > 0 {
				q = args[0]
			}
			params := api.SearchItemsParams{
				Query:      q,
				ItemType:   c.Flags().Lookup("type").Value.String(),
				RarityName: c.Flags().Lookup("rarity").Value.String(),
				WeaponType: c.Flags().Lookup("weapon-type").Value.String(),
				Category:   c.Flags().Lookup("category").Value.String(),
			}
			params.Limit, _ = c.Flags().GetInt("limit")
			params.Offset, _ = c.Flags().GetInt("offset")

			client := newAPIClient()
			resp, err := client.SearchItems(c.Context(), params)
			if err != nil {
				return err
			}

			return renderOutput(renderData{
				data: resp,
				toTable: func() ([]string, [][]string) {
					if len(resp.Items) == 0 {
						return nil, nil
					}
					header := []string{"ID", "Name", "Type", "Rarity", "Wear"}
					rows := make([][]string, len(resp.Items))
					for i, item := range resp.Items {
						rows[i] = []string{
							output.OptionalInt(item.ItemID),
							item.MarketHashName,
							output.Optional(item.ItemType),
							output.Optional(item.RarityName),
							output.Optional(item.WearName),
						}
					}
					return header, rows
				},
			})
		},
	}
	searchCmd.Flags().String("q", "", "Search query")
	searchCmd.Flags().String("type", "", "Filter by item type (weapon, glove, sticker, etc.)")
	searchCmd.Flags().String("rarity", "", "Filter by rarity name")
	searchCmd.Flags().String("weapon-type", "", "Filter by weapon type")
	searchCmd.Flags().String("category", "", "Filter by category")
	searchCmd.Flags().Int("limit", 20, "Maximum results")
	searchCmd.Flags().Int("offset", 0, "Result offset")

	getCmd := &cobra.Command{
		Use:   "get [item-id]",
		Short: "Get item details by ID",
		Args:  cobra.ExactArgs(1),
		Example: `  cs2cap items get 42`,
		RunE: func(c *cobra.Command, args []string) error {
			itemID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			client := newAPIClient()
			item, err := client.GetItem(c.Context(), itemID)
			if err != nil {
				return err
			}

			return renderOutput(renderData{
				data: item,
				toTable: func() ([]string, [][]string) {
					header := []string{"Field", "Value"}
					rows := [][]string{
						{"ID", output.OptionalInt(item.ItemID)},
						{"Name", item.MarketHashName},
						{"Phase", output.Optional(item.Phase)},
						{"Type", output.Optional(item.ItemType)},
						{"Subtype", output.Optional(item.ItemSubtype)},
						{"Weapon", output.Optional(item.WeaponType)},
						{"Base", output.Optional(item.BaseName)},
						{"Skin", output.Optional(item.SkinName)},
						{"Wear", output.Optional(item.WearName)},
						{"Rarity", output.Optional(item.RarityName)},
						{"Collection", output.Optional(item.Collection)},
						{"StatTrak", output.OptionalBool(item.IsStatTrak)},
						{"Souvenir", output.OptionalBool(item.IsSouvenir)},
						{"Min Float", output.OptionalFloat(item.MinFloat)},
						{"Max Float", output.OptionalFloat(item.MaxFloat)},
						{"Supply", output.OptionalInt(item.Supply)},
					}
					return header, rows
				},
			})
		},
	}

	cmd.AddCommand(searchCmd)
	cmd.AddCommand(getCmd)

	return cmd
}
