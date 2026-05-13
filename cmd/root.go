package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cs2cap/cli/internal/api"
	"github.com/cs2cap/cli/internal/config"
	"github.com/cs2cap/cli/internal/output"
)

const welcome = `  ╔═══════════════════════════════════════════════════════╗
  ║               Welcome to CS2Cap CLI                  ║
  ╚═══════════════════════════════════════════════════════╝

  Query CS2 skin prices, bids, sales, and item data
  across 40+ marketplaces from your terminal.
`

type renderData struct {
	data    interface{}
	toTable func() (header []string, rows [][]string)
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cs2cap",
		Short: "CLI for the CS2Cap API",
		Long: `Command-line interface for querying Counter-Strike 2 marketplace data
from the CS2Cap API. Provides unified access to prices, bids, sales,
item catalog, and market analytics across 40+ marketplaces.

Documentation: https://docs.cs2cap.com`,
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
	apiKey := viper.GetString("api_key")
	if apiKey == "" {
		apiKey = promptForAPIKey()
	}
	return api.NewClient(
		apiKey,
		viper.GetString("base_url"),
	)
}

func promptForAPIKey() string {
	fmt.Fprint(os.Stderr, welcome)
	fmt.Fprint(os.Stderr, "\n  You need an API key to get started.\n")
	fmt.Fprint(os.Stderr, "  Get one for free at https://cs2cap.com\n\n")
	fmt.Fprint(os.Stderr, "  Paste your API key and press Enter: ")

	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key == "" {
		fmt.Fprintln(os.Stderr, "\n  No key entered. Run 'cs2cap config init' later to set one up.")
		os.Exit(1)
	}

	cfg := config.Defaults()
	cfg.APIKey = key
	if err := config.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "  Failed to save config: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "  ✓ Saved to ~/.cs2cap.yaml")
	return key
}

func getOutputFormat() (output.Format, error) {
	return output.ParseFormat(viper.GetString("output"))
}

func renderOutput(rd renderData) error {
	format, err := getOutputFormat()
	if err != nil {
		return err
	}

	switch format {
	case output.FormatJSON:
		return output.NewJSONRenderer().Render(os.Stdout, rd.data)
	default:
		header, rows := rd.toTable()
		if header == nil {
			_, err := fmt.Fprintln(os.Stdout, "No results.")
			return err
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
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
}
