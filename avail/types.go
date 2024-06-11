package avail

type ID struct {
	Height   uint64   `json:"blockHeight"`
	ExtIndex uint32   `json:"extIdx"`
	BlobHash [32]byte `json:"blobHash"`
}

type Commitment [32]byte

type Proof = DataProofRPCResponse

type DataProofRPCResponse struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		DataProof `json:"dataProof"`
	} `json:"result"`
}

type DataProof struct {
	Leaf           string   `json:"leaf"`
	LeafIndex      int64    `json:"leafIndex"`
	NumberOfLeaves int64    `json:"numberOfLeaves"`
	Proof          []string `json:"proof"`
	Roots          struct {
		BlobRoot   string `json:"blobRoot"`
		BridgeRoot string `json:"bridgeRoot"`
		DataRoot   string `json:"dataRoot"`
	} `json:"roots"`
}
