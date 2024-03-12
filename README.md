# go-daash

## Prerequisites

You need to have a running Celestia light node to be able to post blobs through `go-daash`. You can find the instructions to run the light node [here](https://docs.celestia.org/developers/node-tutorial).

Copy the `.env.example` file to `.env` and set the `CELESTIA_AUTH_TOKEN` environment variable to the auth token of your Celestia light node.

```bash
cp .env.example .env
```

## Installation/Running

```bash
go build .
go run .
curl --location 'localhost:8080/Celestia' \
    --header 'Content-Type: application/json' \
    --data 'gm'
```
