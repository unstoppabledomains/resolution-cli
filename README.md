# Resolution CLI

Simple CLI tool for resolving .crypto and .zil domains

## Quick Start

### Installation

`go get -u github.com/unstoppabledomains/resolution-cli/resolution`

### Run

`resolution --help`

## Ethereum Provider

If no Ethereum JSON RPC ethereum-provider-url is defined, a default Infura Ethereum Provider Url will be used

Options for defining a ethereumProviderUrlFlag:

- `ETHEREUM_PROVIDER_URL` env variable

- `--ethereum-provider-url` flag

The CLI prioritizes the `--ethereum-provider-url` flag

## Zilliqa Provider

If no Ethereum JSON RPC ethereum-provider-url is defined, a default Zilliqa mainnet Provider Url will be used

Options for defining a ethereumProviderUrlFlag:

- `ZILLIQA_PROVIDER_URL` env variable

- `--zilliqa-provider-url` flag

The CLI prioritizes the `--zilliqa-provider-url` flag

## Resolve a Domain

#### Address

Resolve single address from ticker symbol

`resolution resolve addr ETH -d brad.crypto`

#### IPFS

Resolve ipfs hash

`resolution resolve ipfs-hash -d brad.zil`

#### All

Resolve all known records

`resolution resolve -d brad.zil`

#### Records

Resolve records from exact record keys. See supported records reference [here](https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference)

`resolution resolve records crypto.ETH.adddress crypto.BTC.address -d brad.crypto`

## Development

### Run end-to-end tests
```shell
$ bash run-e2e-test.sh
```
