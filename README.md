# resolution-cli
![Test](https://github.com/unstoppabledomains/resolution-cli/actions/workflows/e2e-test.yml/badge.svg?branch=master)
![Lint](https://github.com/unstoppabledomains/resolution-cli/actions/workflows/lint.yml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/unstoppabledomains/resolution-cli)](https://goreportcard.com/report/github.com/unstoppabledomains/resolution-cli)
[![GoDoc](https://godoc.org/github.com/unstoppabledomains/resolution-cli?status.svg)](https://pkg.go.dev/github.com/unstoppabledomains/resolution-cli)
[![Unstoppable Domains Documentation](https://img.shields.io/badge/docs-unstoppabledomains.com-blue)](https://docs.unstoppabledomains.com/)
[![Get help on Discord](https://img.shields.io/badge/Get%20help%20on-Discord-blueviolet)](https://discord.gg/b6ZVxSZ9Hn)

Simple CLI tool for resolving .crypto and .zil domains

resolution-cli is a tool for interacting with blockchain domain names. It can be used to retrieve [payment addresses](https://unstoppabledomains.com/features#Add-Crypto-Addresses), IPFS hashes for [decentralized websites](https://unstoppabledomains.com/features#Build-Website), DNS records and other [records types](https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference)

resolution-cli is primarily built and maintained by [Unstoppable Domains](https://unstoppabledomains.com/).

resolution-cli supports decentralized domains across three zones:

- Crypto Name Service (CNS)
    - `.crypto`
- Zilliqa Name Service (ZNS)
    - `.zil`

# Quick Start

## Installation

```shell
$ go get -u github.com/unstoppabledomains/resolution-cli/resolution
```

#### Build and install from sources
```shell
$ go build -o ./cli ./resolution
$ mv ./cli /usr/local/bin/resolution
```

## Run

`resolution --help`

### Ethereum Provider

If no Ethereum JSON RPC ethereum-provider-url is defined, a default Infura Ethereum Provider Url will be used

Options for defining a ethereumProviderUrlFlag:

- `ETHEREUM_PROVIDER_URL` env variable

- `--ethereum-provider-url` flag

The CLI prioritizes the `--ethereum-provider-url` flag

### Zilliqa Provider

If no Ethereum JSON RPC ethereum-provider-url is defined, a default Zilliqa mainnet Provider Url will be used

Options for defining a ethereumProviderUrlFlag:

- `ZILLIQA_PROVIDER_URL` env variable

- `--zilliqa-provider-url` flag

The CLI prioritizes the `--zilliqa-provider-url` flag

# Resolve a Domain

## Address

Resolve single address from ticker symbol

`resolution resolve addr ETH -d brad.crypto`

## IPFS

Resolve ipfs hash

`resolution resolve ipfs-hash -d brad.zil`

## All

Resolve all known records

`resolution resolve -d brad.zil`

## Raw records

Resolve records from exact record keys. See supported records reference [here](https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference)

`resolution resolve records crypto.ETH.adddress crypto.BTC.address -d brad.crypto`

# Output format
CLI provides JSON output or quoted string in case of single return value

# Network support
CLI supports Ethereum mainnet and Zilliqa mainnet only.

# Contributions
Contributions are more than welcome. The easiest way to contribute is through GitHub issues and pull requests.

# Development

## Run end-to-end tests
```shell
$ bash run-e2e-test.sh
```

# Free advertising for integrated apps

Once your app has a working Unstoppable Domains integration, [register it here](https://unstoppabledomains.com/app-submission). Registered apps appear on the Unstoppable Domains [homepage](https://unstoppabledomains.com/) and [Applications](https://unstoppabledomains.com/apps) page — putting your app in front of tens of thousands of potential customers per day.

Also, every week we select a newly-integrated app to feature in the Unstoppable Update newsletter. This newsletter is delivered to straight into the inbox of ~100,000 crypto fanatics — all of whom could be new customers to grow your business.

# Get help
[Join our discord community](https://discord.com/invite/b6ZVxSZ9Hn) and ask questions.  