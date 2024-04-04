package daash

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/test"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/celestiada"
	"github.com/stackrlabs/go-daash/eigenda"
)

const CelestiaClientUrl = "http://localhost:26658"

type DALayer string

const (
	Avail    DALayer = "avail"
	Eigen    DALayer = "eigen"
	Celestia DALayer = "celestia"
	Mock     DALayer = "mock"
)

func IsValidDA(layer DALayer) bool {
	for _, validLayer := range []DALayer{Avail, Eigen, Celestia, Mock} {
		if layer == validLayer {
			return true
		}
	}
	return false
}

type DABuilder struct {
	Clients map[DALayer]da.DA
}

func NewDABuilder() *DABuilder {
	return &DABuilder{
		Clients: make(map[DALayer]da.DA),
	}
}

// Initiates a new DAManager with clients from the sepcified DA layers
func (d *DABuilder) InitClients(ctx context.Context, layers []DALayer, availConfigPath string, celestiaAuthToken string) (*DABuilder, error) {
	if len(layers) == 0 {
		return nil, fmt.Errorf("no da layers provided")
	}

	for _, layer := range layers {
		switch layer {
		case Avail:
			var avail da.DA
			var err error
			err = backoff.Retry(func() error {
				avail, err = availda.New(availConfigPath)
				return err //nolint: wrapcheck
			}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
			if err != nil {
				log.Println("❌ Failed to initialise Avail DA client")
				return nil, fmt.Errorf(" Failed to create avail client: %v", err)
			}
			log.Println("🟢 Avail DA client initialised")
			d.Clients[Avail] = avail

		case Celestia:
			if celestiaAuthToken == "" {
				fmt.Println("AUTH_TOKEN is not set")
				return nil, fmt.Errorf("celestia auth token is not set")
			}
			// We use a random pre-set hex string for namespace rn
			namespace := "9cb73e106b03d1050a13"
			celestia, err := celestiada.New(ctx, CelestiaClientUrl, celestiaAuthToken, namespace, -1)
			if err != nil {
				return nil, err
			}
			log.Println("🟢 Celestia DA client initialised")
			d.Clients[Celestia] = celestia

		case Eigen:
			eigen, err := eigenda.New("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
			if err != nil {
				return nil, err
			}
			d.Clients[Eigen] = eigen
			log.Println("🟢 Eigen DA client initialised")

		case Mock:
			d.Clients[Mock] = test.NewDummyDA()
			log.Println("🟢 Mock DA client initialised")

		default:
			return nil, fmt.Errorf("invalid da layer provided: %s", layer)
		}
	}
	return d, nil
}
