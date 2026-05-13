# cs2cap-cli

**Command-line interface for the [CS2Cap API](https://docs.cs2c.app)** — query Counter-Strike 2 marketplace intelligence across 40+ providers from your terminal.

```bash
cs2cap-cli prices list --name "AK-47 | Redline (Field-Tested)"
cs2cap-cli items search --q "Bayonet" --type weapon --output json
cs2cap-cli providers
```

## Features

- **Prices** — query lowest ask listings across all marketplaces
- **Bids** — query highest buy orders across all marketplaces
- **Sales** — recent sales history per item
- **Items** — search/filter the CS2 item catalog
- **Providers** — list all supported marketplace providers
- **Batch lookups** — query multiple items at once by ID or name
- **Multiple output formats** — table (default) or JSON
- **Config file** — persist your API key in `~/.cs2cap.yaml`
- **Environment variable overrides** — `CS2CAP_API_KEY`, `CS2CAP_BASE_URL`, `CS2CAP_OUTPUT`

## Requirements

- Go 1.21+ (to build from source)
- A CS2Cap API key (`sk_live_...` or `sk_test_...`) — sign up at [cs2c.app](https://cs2c.app)

## Installation

### From source

```bash
go install github.com/cs2cap/cli@latest
```

### Build from source

```bash
git clone https://github.com/cs2cap/cli.git
cd cli
make build
./cs2cap-cli --help
```

The binary is statically compiled with no runtime dependencies.

## Quick Start

```bash
# Set your API key (or use ~/.cs2cap.yaml)
export CS2CAP_API_KEY=sk_live_your_key_here

# List prices for an item
cs2cap-cli prices list --name "AK-47 | Redline (Field-Tested)"

# Search for items
cs2cap-cli items search --q "Bayonet" --limit 5

# Get item details
cs2cap-cli items get 42

# List all providers
cs2cap-cli providers
```

## Configuration

cs2cap-cli resolves settings in this order (last wins):

1. CLI flags
2. Environment variables
3. Config file (`~/.cs2cap.yaml`)
4. Built-in defaults

### Config file

```bash
# Interactive setup (writes ~/.cs2cap.yaml)
cs2cap-cli config init

# View current config (key masked)
cs2cap-cli config show
```

Example `~/.cs2cap.yaml`:

```yaml
api_key: sk_live_your_key_here
base_url: https://api.cs2c.app
output: table
```

### Environment variables

| Variable | Description | Default |
| --- | --- | --- |
| `CS2CAP_API_KEY` | API key | — |
| `CS2CAP_BASE_URL` | API base URL | `https://api.cs2c.app` |
| `CS2CAP_OUTPUT` | Output format | `table` |

### Global flags

| Flag | Short | Description |
| --- | --- | --- |
| `--api-key` | `-k` | API key (overrides config & env) |
| `--base-url` | | API base URL |
| `--output` | `-o` | Output format: `table` or `json` |
| `--help` | `-h` | Help for any command |

## Usage

### Prices

List current lowest ask prices across providers:

```bash
cs2cap-cli prices list --name "AK-47 | Redline (Field-Tested)"
cs2cap-cli prices list --item-id 1234 --providers steam --providers buff163
cs2cap-cli prices list --name "★ Bayonet | Doppler" --phase ruby --currency EUR
cs2cap-cli prices list --output json
```

Batch price lookup by item IDs or names:

```bash
cs2cap-cli prices batch --items 1,2,3
cs2cap-cli prices batch --names "AK-47 | Redline (FT)","★ Bayonet | Doppler"
```

### Items

Search/filter the item catalog:

```bash
cs2cap-cli items search --q "AK-47"
cs2cap-cli items search --type weapon --rarity "Covert"
cs2cap-cli items search --type sticker --limit 50
```

Get full item details by ID:

```bash
cs2cap-cli items get 42
```

### Bids

List current highest buy orders:

```bash
cs2cap-cli bids list --name "AK-47 | Redline (Field-Tested)"
cs2cap-cli bids list --item-id 1234 --providers steam --providers buff163
cs2cap-cli bids batch --items 1,2,3
```

### Sales

View recent sales history:

```bash
cs2cap-cli sales list --name "AK-47 | Redline (Field-Tested)"
cs2cap-cli sales list --item-id 1234 --providers steam --limit 10
```

### Providers

List all supported marketplace providers:

```bash
cs2cap-cli providers
cs2cap-cli providers --output json
```

## Output Formats

### Table (default)

Human-readable aligned columns with tab-separated values.

### JSON

Machine-readable JSON output using `json.MarshalIndent`:

```bash
cs2cap-cli prices list --name "AK-47 | Redline (FT)" --output json
```

## Price Formatting

All prices from the CS2Cap API are returned in **minor units** (cents). cs2cap-cli automatically divides by 100 for display:

```
$1,234.56
```

## Building

```bash
git clone https://github.com/cs2cap/cli.git
cd cli
make build      # produces ./cs2cap-cli
make lint       # run go vet
make test       # run tests
make clean      # remove binary
make tidy       # clean go.mod dependencies
```

## Dependencies

cs2cap-cli has minimal dependencies:

| Package | Purpose |
| --- | --- |
| [spf13/cobra](https://github.com/spf13/cobra) | CLI framework |
| [spf13/viper](https://github.com/spf13/viper) | Configuration management |
| [go-yaml/yaml.v3](https://gopkg.in/yaml.v3) | YAML config file parsing |
| Go standard library | HTTP client, tabular output (`text/tabwriter`), JSON encoding |

## API Documentation

Full CS2Cap API documentation is available at [docs.cs2c.app](https://docs.cs2c.app).

## License

MIT
