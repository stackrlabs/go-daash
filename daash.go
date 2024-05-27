package daash

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/stackrlabs/go-daash/avail"
	"github.com/stackrlabs/go-daash/celestia"
	"github.com/stackrlabs/go-daash/da"
	"github.com/stackrlabs/go-daash/eigen"
	"github.com/stackrlabs/go-daash/mock"
)

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

type ClientBuilder struct {
	Clients map[DALayer]da.Client
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		Clients: make(map[DALayer]da.Client),
	}
}

// Initiates a new DAManager with clients from the sepcified DA layers
func (d *ClientBuilder) InitClients(ctx context.Context, layers []DALayer, availConfigPath string, celestiaAuthToken string, celestiaLightClientUrl string) (*ClientBuilder, error) {
	if len(layers) == 0 {
		return nil, fmt.Errorf("no da layers provided")
	}

	for _, layer := range layers {
		switch layer {
		case Avail:
			var availClient da.Client
			var err error
			err = backoff.Retry(func() error {
				availClient, err = avail.NewClient(availConfigPath)
				return err //nolint: wrapcheck
			}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
			if err != nil {
				log.Println("❌ Failed to initialise Avail DA client")
				return nil, fmt.Errorf(" Failed to create avail client: %v", err)
			}
			log.Println("🟢 Avail DA client initialised")
			d.Clients[Avail] = availClient

		case Celestia:
			if celestiaAuthToken == "" {
				fmt.Println("AUTH_TOKEN is not set")
				return nil, fmt.Errorf("celestia auth token is not set")
			}
			// We use a random pre-set hex string for namespace rn
			namespace := "9cb73e106b03d1050a13"
			celestia, err := celestia.NewClient(ctx, celestiaLightClientUrl, celestiaAuthToken, namespace, -1)
			if err != nil {
				return nil, err
			}
			log.Println("🟢 Celestia DA client initialised")
			d.Clients[Celestia] = celestia

		case Eigen:
			eigen, err := eigen.NewClient("disperser-goerli.eigenda.xyz:443", time.Second*90, time.Second*5)
			if err != nil {
				return nil, err
			}
			d.Clients[Eigen] = eigen
			log.Println("🟢 Eigen DA client initialised")

		case Mock:
			d.Clients[Mock] = mock.NewDummyDA()
			log.Println("🟢 Mock DA client initialised")

		default:
			return nil, fmt.Errorf("invalid da layer provided: %s", layer)
		}
	}
	return d, nil
}

func GetHumanReadableID(id da.ID, daLayer DALayer) any {
	switch daLayer {
	case Avail:
		availID, ok := id.(avail.ID)
		if !ok {
			return ""
		}
		return availID
	case Celestia:
		id, ok := id.(celestia.ID)
		if !ok {
			return ""
		}
		return struct {
			BlockHeight uint64        `json:"blockHeight"`
			TxHash      string        `json:"txHash"`
			Commitment  da.Commitment `json:"commitment"`
		}{
			BlockHeight: id.Height,
			TxHash:      hex.EncodeToString(id.TxHash),
			Commitment:  id.Commitment,
		}
	default:
		return ""
	}
}

func GetExplorerLink(client da.Client, ids []da.ID) (string, error) {
	switch daClient := client.(type) {
	case *celestia.Client:
		id, ok := ids[0].(celestia.ID)
		if !ok {
			return "", fmt.Errorf("invalid ID")
		}
		return fmt.Sprintf("https://mocha-4.celenium.io/tx/%s", hex.EncodeToString(id.TxHash)), nil
	case *avail.Client:
		ext, err := daClient.GetExtrinsic(ids[0])
		if err != nil {
			return "", err
		}
		extBytes, err := json.Marshal(ext)
		if err != nil {
			return "", err
		}
		// Strip string of any leading or following quotes
		extString := strings.Trim(string(extBytes), "\"")
		fmt.Println(extString)
		return fmt.Sprintf("https://goldberg.avail.tools/#/extrinsics/decode/%s", extString), nil
	default:
		return "", nil
	}
}
