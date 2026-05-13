# cs2cap

**CLI for the [CS2Cap API](https://docs.cs2cap.com)** — query CS2 skin prices, bids, sales, and item data across 40+ marketplaces from your terminal.

```bash
cs2cap prices list --name "AK-47 | Redline FT"
cs2cap items search --q "Bayonet" --type weapon --output json
cs2cap providers
```

## Features

- **Prices** — lowest ask listings across all marketplaces
- **Bids** — highest buy orders
- **Sales** — recent sales history per item
- **Items** — search/filter the CS2 item catalog
- **Providers** — list all supported marketplace providers
- **Batch lookups** — query multiple items at once by ID or name
- **Wear shortcuts** — `FT`, `MW`, `FN`, `WW`, `BS` auto-expand in item names
- **Multiple output formats** — table (default) or JSON
- **Config file** — persist your API key in `~/.cs2cap.yaml`
- **Environment variable overrides** — `CS2CAP_API_KEY`, `CS2CAP_BASE_URL`, `CS2CAP_OUTPUT`
- **Onboarding** — first-run setup guide when no API key is configured

## Requirements

- Go 1.21+ (to build from source)
- A CS2Cap API key (`sk_live_...`) — sign up at [cs2cap.com](https://cs2cap.com)

## Installation

### From source

```bash
go install github.com/cs2cap/cli/cmd/cs2cap@latest
```

This produces a `cs2cap` binary in your `$GOPATH/bin` or `$HOME/go/bin`.

### Build from source

```bash
git clone https://github.com/cs2cap/cli.git
cd cli
make build
./cs2cap --help
```

## Quick Start

The first time you run a command without an API key, cs2cap shows an onboarding guide:

```bash
cs2cap prices list
# → Welcome to CS2Cap CLI — guides you through setup
```

Once you have a key:

```bash
# Set your API key (or use cs2cap config init)
export CS2CAP_API_KEY=sk_live_your_key_here

# List prices for an item (wear shortcut auto-expands)
cs2cap prices list --name "AK-47 | Redline FT"

# Search for items
cs2cap items search --q "Bayonet" --limit 5

# Get item details
cs2cap items get 42

# List all providers
cs2cap providers
```

## Configuration

cs2cap resolves settings in this order (last wins):

1. CLI flags
2. Environment variables
3. Config file (`~/.cs2cap.yaml`)
4. Built-in defaults

### Config file

```bash
# Interactive setup (writes ~/.cs2cap.yaml)
cs2cap config init

# View current config (key masked)
cs2cap config show
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

## Wear Shortcuts

Wear abbreviations in `--name` and `--names` flags auto-expand:

| Shortcut | Expanded |
| --- | --- |
| `FN` | `Factory New` |
| `MW` | `Minimal Wear` |
| `FT` | `Field-Tested` |
| `WW` | `Well-Worn` |
| `BS` | `Battle-Scarred` |

Both `FT` and `(FT)` work — both expand to `(Field-Tested)`.

## Commands

### `prices list`

Query current lowest ask prices across marketplaces.

**Usage:** `cs2cap prices list [name] [flags]`
**API:** `GET /v1/prices`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--name` | `string` | `""` | Exact market hash name (or pass as first arg) |
| `--item-id` | `int` | `0` | Filter by item ID |
| `--phase` | `string` | `""` | Filter by Doppler phase (`ruby`, `sapphire`, `emerald`, `black_pearl`) |
| `--providers` | `[]string` | `nil` | Filter by provider keys (repeat flag: `--providers steam --providers buff163`) |
| `--currency` | `string` | `"USD"` | Quote currency code |
| `--limit` | `int` | `20` | Maximum results |
| `--offset` | `int` | `0` | Result offset for pagination |

**Examples:**

```bash
cs2cap prices list "AK-47 | Redline FT"
cs2cap prices list --item-id 1234 --providers steam --providers buff163
cs2cap prices list "★ Bayonet | Doppler" --phase ruby --currency EUR
cs2cap prices list --output json
```

### `prices batch` (Starter — $19/mo)

Batch price lookup for multiple items at once.

**Usage:** `cs2cap prices batch [flags]`
**API:** `POST /v1/prices/batch`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--items` | `[]int` | `nil` | Comma-separated item IDs (`--items 1,2,3`) |
| `--names` | `[]string` | `nil` | Comma-separated market hash names (`--names "AK-47 | Redline FT","★ Bayonet | Doppler"`) |

**Examples:**

```bash
cs2cap prices batch --items 1,2,3
cs2cap prices batch --names "AK-47 | Redline FT","★ Bayonet | Doppler"
```

---

### `bids list` (Starter — $19/mo)

Query current highest buy orders across marketplaces.

**Usage:** `cs2cap bids list [name] [flags]`
**API:** `GET /v1/bids`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--name` | `string` | `""` | Exact market hash name (or pass as first arg) |
| `--item-id` | `int` | `0` | Filter by item ID |
| `--phase` | `string` | `""` | Filter by Doppler phase |
| `--providers` | `[]string` | `nil` | Filter by provider keys (repeat flag) |
| `--currency` | `string` | `"USD"` | Quote currency code |
| `--limit` | `int` | `20` | Maximum results |
| `--offset` | `int` | `0` | Result offset for pagination |

**Examples:**

```bash
cs2cap bids list "AK-47 | Redline FT"
cs2cap bids list --item-id 1234 --providers steam --providers buff163
```

### `bids batch` (Starter — $19/mo)

Batch bid lookup for multiple items at once.

**Usage:** `cs2cap bids batch [flags]`
**API:** `POST /v1/bids/batch`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--items` | `[]int` | `nil` | Comma-separated item IDs (`--items 1,2,3`) |
| `--names` | `[]string` | `nil` | Comma-separated market hash names |

**Examples:**

```bash
cs2cap bids batch --items 1,2,3
cs2cap bids batch --names "AK-47 | Redline FT","★ Bayonet | Doppler"
```

---

### `sales list` (Pro — $79/mo)

Query recent sales history for an item.

**Usage:** `cs2cap sales list [name] [flags]`
**API:** `GET /v1/sales`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--name` | `string` | `""` | Exact market hash name (or pass as first arg) |
| `--item-id` | `int` | `0` | Filter by item ID |
| `--providers` | `[]string` | `nil` | Filter by sales-capable provider keys (repeat flag) |
| `--limit` | `int` | `20` | Maximum results (capped at 50) |

**Examples:**

```bash
cs2cap sales list "AK-47 | Redline FT"
cs2cap sales list --item-id 1234 --providers steam --limit 10
```

---

### `items search`

Search or filter the CS2 item catalog.

**Usage:** `cs2cap items search [query] [flags]`
**API:** `GET /v1/items`

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--q` | `string` | `""` | Search query |
| `--type` | `string` | `""` | Filter by item type (`weapon`, `glove`, `sticker`, `collectible`, `agent`, `key`, `tool`, `music_kit`, `patch`, `graffiti`, `name_tag`) |
| `--rarity` | `string` | `""` | Filter by rarity name |
| `--weapon-type` | `string` | `""` | Filter by weapon type |
| `--category` | `string` | `""` | Filter by category |
| `--limit` | `int` | `20` | Maximum results |
| `--offset` | `int` | `0` | Result offset for pagination |

**Examples:**

```bash
cs2cap items search "AK-47"
cs2cap items search --type weapon --rarity "Covert"
cs2cap items search --type sticker --limit 50
```

### `items get`

Get detailed information for a single item by ID.

**Usage:** `cs2cap items get <item-id>`
**API:** `GET /v1/items/{id}`

Returns all available fields including phase, wear range, float values, StatTrak/Souvenir flags, collection, and supply data.

**Example:**

```bash
cs2cap items get 42
```

---

### `providers`

List all supported marketplace providers and their status.

**Usage:** `cs2cap providers`
**API:** `GET /v1/providers`

No command-specific flags. Shows each provider's key, code, market type, default currency, supported features (bids, sales), and health status.

**Examples:**

```bash
cs2cap providers
cs2cap providers --output json
```

---

### `config init`

Interactive setup for `~/.cs2cap.yaml`. Prompts for your API key and saves the configuration file.

**Usage:** `cs2cap config init`

Prompts for confirmation before overwriting an existing config file.

### `config show`

Display the currently active configuration with the API key masked.

**Usage:** `cs2cap config show`

Shows config file path, masked API key, base URL, and output format.

## Output Formats

### Table (default)

Aligned columns with tab-separated values, rendered to the terminal.

### JSON

Machine-readable output:

```bash
cs2cap prices list --name "AK-47 | Redline FT" --output json
```

## Price Formatting

All prices from the CS2Cap API are returned in **minor units** (cents). cs2cap automatically divides by 100 for display:

```
$1,234.56
```

## Building

```bash
make build      # produces ./cs2cap
make lint       # run go vet
make test       # run tests
make clean      # remove binary
make tidy       # clean go.mod dependencies
```

## Dependencies

cs2cap has minimal dependencies:

| Package | Purpose |
| --- | --- |
| [spf13/cobra](https://github.com/spf13/cobra) | CLI framework |
| [spf13/viper](https://github.com/spf13/viper) | Configuration management |
| [go-yaml/yaml.v3](https://gopkg.in/yaml.v3) | YAML config file parsing |
| Go standard library | HTTP client, tabular output (`text/tabwriter`), JSON encoding |

## API Documentation

Full CS2Cap API documentation is available at [https://docs.cs2cap.com](https://docs.cs2cap.com).

## License

MIT
