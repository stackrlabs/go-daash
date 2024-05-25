package main

// Constants
const (
	EigenDaRpcUrl  = "disperser-goerli.eigenda.xyz:443"
	celestiaRpcUrl = "https://celestia-mocha-rpc.publicnode.com:443"
)

var chainMetadata = map[string]map[string]string{
	"sepolia": {
		"rpcUrl":                    "https://sepolia.drpc.org",
		"chainID":                   "1115511",
		"blobstreamverifierAddress": "0x1Bf80E9b8d21ddCCE11b221E1a23781FEb58EB19", // Contract deployed here: https://sepolia.etherscan.io/address/0x1bf80e9b8d21ddcce11b221e1a23781feb58eb19
		"blobstreamxAddress":        "0xf0c6429ebab2e7dc6e05dafb61128be21f13cb1e", // BlobstreamX contract deployed on Sepolia
		"vectorXAddress":            "0xe542dB219a7e2b29C7AEaEAce242c9a2Cd528F96",
		"availBridgeAddress":        "0x967F7DdC4ec508462231849AE81eeaa68Ad01389", // Avail bridge deployed on Sepolia
		"vectorVerifierAddress":     "0x6B26173C8afF316919542df8dA5A57888e398ee1", // Custom Vector verifier contract deployed on Sepolia
	},
}
