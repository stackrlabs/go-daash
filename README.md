# go-daash 🏃‍♂️

![cover](./assets/cover.png)

Unified interaction API for Data Availability chains initially designed to work with Vulcan (Stackr's Verification Layer)

This unified API for data availability services provides a scalable, secure, and efficient way for Vulcan to interact with different DA networks. It aims to standardize data access patterns, reduce complexity, and enhance the overall experience of using data availability services.

## Demo

Checkout go-daash in action -> [BlobDaddy 👨🏻](https://blob-daddy.stf.xyz/) 

## Motivation

As an un-opinionated/agnostic framework Stackr should support all different kind of DA layers and provide a smooth abstracted experience across the different DA layers.

Each DA layer has a different RPC, a different API and different Data formats. Each service also requires holding their native tokens and there are many other nuances to these as well including security models etc.

Every micro-rollup should be easily able to switch between different DA networks with minimal changes to their codebase.

## Supported Networks

1. Avail
2. EigenDA
3. Celestia
4. (planned) NearDA
5. (planned) EIP-4844

## Features

1. Runs with Vulcan but can be run as a standalone service
2. Separate endpoints for each DA network
3. Can post blobs to the respected DA network

## Planned Features

1. Blob archival management
2. Backup DA posting if posting fails on one DA
3. Blob chunking if max blob size is reached
4. Payment abstraction: Pay conveniently in ETH or USD and let DaasH take care of token management

## Development

### Steps to run

```bash
git clone https://github.com/stackrlabs/go-daash # clone the repository
cd go-daash/cmd/blob-server
go run . # run the server
curl --location 'localhost:8080/eigen' \
--header 'Content-Type: application/json' \
--data 'gm' # DAash away!
```

### Additional requirements

Avail and EigenDA run via external RPC providers while Celestia requires a running light node to be able to post blobs through `go-daash`.

For Avail, you need to provide the path to a JSON config file with the following structure:

```json
{
  "seed": "<your seed phrases>",
  "api_url": "wss://goldberg.avail.tools:443/ws",
  "app_id": 1,
  "timeout": 600
}
```

For Celestia, you can find the instructions to run the light node [here](https://docs.celestia.org/developers/node-tutorial).
You need an auth token for your Celestia light node. Copy the `.env.example` file to `.env` and set the `CELESTIA_AUTH_TOKEN` environment variable to the auth token.

```bash
cp .env.example .env
```

## Built with
- `go-da` and `celestia-da` libraries from [Rollkit](https://github.com/rollkit/go-da)
- `blobstreamx-example` repo from [CryptoKass](https://github.com/CryptoKass/blobstreamx-example)


