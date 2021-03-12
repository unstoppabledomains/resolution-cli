# Resolution CLI

Simple CLI tool for resolving .crypto and .zil domains

## Quick Start

### Installation

`go get -u github.com/unstoppabledomains/resolution-cli/resolution`

### Run

`resolution --help`

## Ethereum Provider

If no Ethereum JSON RPC provider is defined, a default Infura provider will be used

Options for defining a provider:

- `RESOLUTION_PROVIDER` env variable

- `--provider` flag

The CLI prioritizes the `--provider` flag

## Flags

`--provider` or `-p`

- Provider to use

`--domain` or `-d`

- Domain to use. This flag is required for the resolve command

## Resolve a domain

#### Address

Resolves single address from ticker symbol

`resolution resolve addr ETH -d brad.crypto`

#### IPFS

Resolves ipfs hash

`resolution resolve ipfs -d brad.zil`

#### All

Resolves all records

`resolution resolve records -d brad.zil`

#### Records

Resolves records from exact record keys. See supported records reference [here](https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference)

`resolution resolve records crypto.ETH.adddress crypto.BTC.address -d brad.crypto`
