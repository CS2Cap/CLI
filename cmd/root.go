package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/output"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cs2cap",
		Short: "CLI for the CS2Cap API",
		Long: `Command-line interface for querying Counter-Strike 2 marketplace data
from the CS2Cap API. Provides unified access to prices, bids, sales,
item catalog, and market analytics across 40+ marketplaces.

Documentation: https://docs.cs2c.app`,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			viper.SetConfigName(".cs2cap")
			viper.SetConfigType("yaml")
			viper.AddConfigPath("$HOME")
			if err := viper.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
					return err
				}
			}
			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
	}

	cmd.PersistentFlags().StringP("api-key", "k", "", "API key (env: CS2CAP_API_KEY)")
	cmd.PersistentFlags().String("base-url", "https://api.cs2c.app", "API base URL (env: CS2CAP_BASE_URL)")
	cmd.PersistentFlags().StringP("output", "o", "table", "Output format: table|json (env: CS2CAP_OUTPUT)")

	viper.BindPFlag("api_key", cmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("base_url", cmd.PersistentFlags().Lookup("base-url"))
	viper.BindPFlag("output", cmd.PersistentFlags().Lookup("output"))
	viper.SetEnvPrefix("CS2CAP")
	viper.AutomaticEnv()

	return cmd
}

func Execute() {
	root := newRootCmd()

	root.AddCommand(newConfigCmd())
	root.AddCommand(newPricesCmd())
	root.AddCommand(newItemsCmd())
	root.AddCommand(newBidsCmd())
	root.AddCommand(newSalesCmd())
	root.AddCommand(newProvidersCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func newAPIClient() *api.Client {
	return api.NewClient(
		viper.GetString("api_key"),
		viper.GetString("base_url"),
	)
}

func getOutputFormat() (output.Format, error) {
	return output.ParseFormat(viper.GetString("output"))
}

func renderOutput(v interface{}) error {
	format, err := getOutputFormat()
	if err != nil {
		return err
	}

	var r output.Renderer
	switch format {
	case output.FormatJSON:
		r = output.NewJSONRenderer()
	default:
		r = output.NewTableRenderer()
	}

	return r.Render(os.Stdout, v)
}
