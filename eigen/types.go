package eigen

import grpcdisperser "github.com/Layr-Labs/eigenda/api/grpc/disperser"

type ID struct {
	BlobInfo  *grpcdisperser.BlobInfo
	RequestID string
}
