# go-daash 🏃‍♂️

Unified interaction API for Data Availibility chains initially designed to work with Vulcan (Stackr's Verification Layer)

This unified API for data availability services provides a scalable, secure, and efficient way for Vulcan to interact with different DA networks. It aims to standardize data access patterns, reduce complexity, and enhance the overall experience of using data availability services.

## Motivation

As an un-opinionated/agnostic framework Stackr should support all different kind of DA layers and provide a smooth abstracted experience across the different DA layers.

Each DA layer has a different RPC, a different API and different Data formats. Each service also requires holding their native tokens and there are many other nuances to these as well including security models etc.

Each micro-rollup should be easily able to switch

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

1. Clone the repository
2.

### Additional requirements

Avail and EigenDA runs via external RPC providers however Celestia requires a running light node to be able to post blobs through `go-daash`.

You can find the instructions to run the light node [here](https://docs.celestia.org/developers/node-tutorial).
Also export auth token as environment variable:

```bash
export CELESTIA_AUTH_TOKEN=<your_auth_token>
```

## Installation/Running

```bash
go build .
go run .
curl --location 'localhost:8080/Celestia' \
--header 'Content-Type: application/json' \
--data 'gm'
```
