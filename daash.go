package daash

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/eigenda"
)

type DAType string

const (
	Avail DAType = "avail"
	Eigen DAType = "eigen"
)

type DAManager struct {
	Clients map[DAType]da.DA
}

// Initialize all the DA clients
func (d *DAManager) Init(availConfigPath string) error {
	var err error
	// Initialize Avail
	var avail da.DA
	err = backoff.Retry(func() error {
		avail, err = availda.New(availConfigPath)
		return err //nolint: wrapcheck
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
	if err != nil {
		return err
	}
	d.Clients[Avail] = avail

	// Initialize Eigen
	eigen, err := eigenda.New("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
	if err != nil {
		return err
	}
	d.Clients[Eigen] = eigen

	// TODO: Initialize Celestia

	return nil
}
