package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	DefaultBaseURL = "https://api.cs2c.app"
	DefaultOutput  = "table"
)

type Config struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
	Output  string `yaml:"output"`
}

func Defaults() *Config {
	return &Config{
		BaseURL: DefaultBaseURL,
		Output:  DefaultOutput,
	}
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cs2cap.yaml"), nil
}

func Load() (*Config, error) {
	cfg := Defaults()

	path, err := configPath()
	if err != nil {
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func Save(cfg *Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Exists() bool {
	path, err := configPath()
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}
