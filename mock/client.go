package mock

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"sync"

	"github.com/stackrlabs/go-daash/da"
)

// DefaultMaxBlobSize is the default max blob size
const DefaultMaxBlobSize = 64 * 64 * 482

// DummyDA is a simple implementation of in-memory DA. Not production ready! Intended only for testing!
//
// Data is stored in a map, where key is a serialized sequence number. This key is returned as ID.
// Commitments are simply hashes, and proofs are ED25519 signatures.
type DummyDA struct {
	mu          *sync.Mutex // protects data and height
	data        map[uint64][]kvp
	maxBlobSize uint64
	height      uint64
	privKey     ed25519.PrivateKey
	pubKey      ed25519.PublicKey
}

type kvp struct {
	key, value []byte
}

// NewDummyDA create new instance of DummyDA
func NewDummyDA(opts ...func(*DummyDA) *DummyDA) *DummyDA {
	da := &DummyDA{
		mu:          new(sync.Mutex),
		data:        make(map[uint64][]kvp),
		maxBlobSize: DefaultMaxBlobSize,
	}
	for _, f := range opts {
		da = f(da)
	}
	da.pubKey, da.privKey, _ = ed25519.GenerateKey(rand.Reader)
	return da
}

// MaxBlobSize returns the max blob size in bytes.
func (d *DummyDA) MaxBlobSize(ctx context.Context) (uint64, error) {
	return d.maxBlobSize, nil
}

// Get returns Blob for given ID.
func (d *DummyDA) Get(ctx context.Context, id da.ID) (da.Blob, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var blob da.Blob
	mockID, ok := id.(ID)
	if !ok {
		return nil, errors.New("invalid ID")
	}
	if len(mockID) < 8 {
		return nil, errors.New("invalid ID")
	}
	height := binary.LittleEndian.Uint64(mockID)
	found := false
	for j := 0; !found && j < len(d.data[height]); j++ {
		if bytes.Equal(d.data[height][j].key, mockID) {
			blob = d.data[height][j].value
			found = true
		}
	}
	if !found {
		return nil, errors.New("no blob for given ID")
	}
	return blob, nil
}

// Commit returns cryptographic Commitments for given blobs.
func (d *DummyDA) Commit(ctx context.Context, blob da.Blob) (da.Commitment, error) {
	return d.getHash(blob), nil
}

// Submit stores blobs in DA layer.
func (d *DummyDA) Submit(ctx context.Context, blob da.Blob, gasPrice float64) (da.ID, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.height += 1
	id := append(d.nextID(), d.getHash(blob)...)
	d.data[d.height] = append(d.data[d.height], kvp{id, blob})

	return id, nil
}

// GetProof returns inclusion Proofs for all Blobs located in DA at given height.
func (d *DummyDA) GetProof(ctx context.Context, id da.ID) (da.Proof, error) {
	blob, err := d.Get(ctx, id)

	d.mu.Lock()
	defer d.mu.Unlock()
	if err != nil {
		return nil, err
	}
	proof := d.getProof(id.(ID), blob)
	return proof, nil
}

// Validate checks the Proofs for given IDs.
func (d *DummyDA) Validate(ctx context.Context, id da.ID, proof da.Proof) (bool, error) {
	result := ed25519.Verify(d.pubKey, id.(ID)[8:], proof.([]byte))

	return result, nil
}

func (d *DummyDA) nextID() []byte {
	return d.getID(d.height)
}

type ID = []byte

func (d *DummyDA) getID(cnt uint64) []byte {
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, cnt)
	return id
}

func (d *DummyDA) getHash(blob []byte) []byte {
	sha := sha256.Sum256(blob)
	return sha[:]
}

func (d *DummyDA) getProof(id ID, blob []byte) []byte {
	sign, _ := d.privKey.Sign(rand.Reader, d.getHash(blob), &ed25519.Options{})
	return sign
}
