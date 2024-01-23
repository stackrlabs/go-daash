# go-daash

## Prerequisites
You need to have a running Celestia light node to be able to post blobs through `go-daash`. You can find the instructions to run the light node [here](https://docs.celestia.org/developers/node-tutorial).
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