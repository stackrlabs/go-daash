package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stackrlabs/go-daash"
)

type Job struct {
	Data   []byte
	Layer  daash.DALayer
	ID     string
	Status map[string]any // Human-readable job status
}

func generateJobID() string {
	b := make([]byte, 16) // Generates a 128-bit (16 bytes) random hex string
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("error generating random hex string: %v", err)
	}
	randomHexString := hex.EncodeToString(b)
	return randomHexString
}

func run(ctx context.Context, b *BlobServer, job Job) {
	var jobStatus map[string]any
	ids, proofs, err := postToDA(ctx, job.Data, b.Daasher.Clients[job.Layer])
	if err != nil {
		jobStatus = gin.H{
			"status": "failed",
			"error":  err,
		}
	} else {
		successLink, err := daash.GetExplorerLink(b.Daasher.Clients[job.Layer], ids)
		if err != nil {
			log.Fatalf("cannot get explorer link: %v", err)
		}
		jobStatus = gin.H{
			"status": "Blob daashed and posted to " + string(job.Layer) + " 🏃",
			"ids":    ids,
			"proofs": proofs,
			"link":   successLink,
		}
	}
	b.Lock()
	b.Jobs[job.ID] = Job{Data: job.Data, Layer: job.Layer, ID: job.ID, Status: jobStatus}
	b.Unlock()
}
