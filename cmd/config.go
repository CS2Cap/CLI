package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cs2cap/cli/internal/config"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Create ~/.cs2cap.yaml configuration file",
		RunE: func(c *cobra.Command, args []string) error {
			if config.Exists() {
				fmt.Fprint(os.Stderr, "Config file already exists at ~/.cs2cap.yaml. Overwrite? [y/N] ")
				var resp string
				fmt.Scanln(&resp)
				if resp != "y" && resp != "Y" {
					fmt.Println("Aborted.")
					return nil
				}
			}

			fmt.Print("API key (sk_live_...): ")
			var apiKey string
			fmt.Scanln(&apiKey)

			cfg := config.Defaults()
			cfg.APIKey = apiKey

			if err := config.Save(cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Println("Configuration saved to ~/.cs2cap.yaml")
			return nil
		},
	})

	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(c *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}

			masked := "not set"
			if cfg.APIKey != "" {
				if len(cfg.APIKey) > 8 {
					masked = cfg.APIKey[:8] + "..." + cfg.APIKey[len(cfg.APIKey)-4:]
				} else {
					masked = "****"
				}
			}

			fmt.Printf("Config file:  ~/.cs2cap.yaml%s\n", map[bool]string{true: " (exists)", false: " (not found)"}[config.Exists()])
			fmt.Printf("API key:      %s\n", masked)
			fmt.Printf("Base URL:     %s\n", cfg.BaseURL)
			fmt.Printf("Output:       %s\n", cfg.Output)
			return nil
		},
	}
	cmd.AddCommand(showCmd)

	return cmd
}
