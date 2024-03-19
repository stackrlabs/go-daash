package daash

import (
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/test"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/eigenda"
)

type DALayer string

const (
	Avail DALayer = "avail"
	Eigen DALayer = "eigen"
	Mock  DALayer = "mock"
)

func IsValidDA(layer DALayer) bool {
	for _, validLayer := range []DALayer{Avail, Eigen, Mock} {
		if layer == validLayer {
			return true
		}
	}
	return false
}

type DAManager struct {
	Clients map[DALayer]da.DA
}

// Initiates a new DAManager with clients from the sepcified DA layers
func NewDAManager(layers []DALayer, availConfigPath string) (*DAManager, error) {
	if len(layers) == 0 {
		return nil, fmt.Errorf("no da layers provided")
	}

	d := &DAManager{}
	d.Clients = make(map[DALayer]da.DA)
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
				return nil, err
			}
			log.Println("🟢 Avail DA client initialised")
			d.Clients[Avail] = avail
		case Eigen:
			eigen, err := eigenda.New("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
			if err != nil {
				return nil, err
			}
			d.Clients[Eigen] = eigen
		case Mock:
			d.Clients[Mock] = test.NewDummyDA()
			log.Println("🟢 Mock DA client initialised")
		default:
			return nil, fmt.Errorf("invalid da layer provided: %s", layer)
		}
	}
	return d, nil
}
