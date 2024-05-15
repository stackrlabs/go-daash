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
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/test"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/celestiada"
	"github.com/stackrlabs/go-daash/eigenda"
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

type DABuilder struct {
	Clients map[DALayer]da.DA
}

func NewDABuilder() *DABuilder {
	return &DABuilder{
		Clients: make(map[DALayer]da.DA),
	}
}

// Initiates a new DAManager with clients from the sepcified DA layers
func (d *DABuilder) InitClients(ctx context.Context, layers []DALayer, availConfigPath string, celestiaAuthToken string, celestiaLightClientUrl string) (*DABuilder, error) {
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
			celestia, err := celestiada.New(ctx, celestiaLightClientUrl, celestiaAuthToken, namespace, -1)
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

func GetHumanReadableID(id da.ID, daLayer DALayer) any {
	switch daLayer {
	case Avail:
		blockHeight, extIdx := availda.SplitID(id)
		return struct {
			BlockHeight uint32 `json:"blockHeight"`
			ExtIdx      uint32 `json:"extIdx"`
		}{
			BlockHeight: blockHeight,
			ExtIdx:      extIdx,
		}
	case Celestia:
		blockHeight, txHash, commitment := celestiada.SplitID(id)
		return struct {
			BlockHeight uint64        `json:"blockHeight"`
			TxHash      string        `json:"txHash"`
			Commitment  da.Commitment `json:"commitment"`
		}{
			BlockHeight: blockHeight,
			TxHash:      hex.EncodeToString(txHash),
			Commitment:  commitment,
		}
	default:
		return ""
	}
}

func GetExplorerLink(client da.DA, ids []da.ID) (string, error) {
	switch daClient := client.(type) {
	case *celestiada.DAClient:
		_, txHash, _ := celestiada.SplitID(ids[0])
		return fmt.Sprintf("https://mocha-4.celenium.io/tx/%s", hex.EncodeToString(txHash)), nil
	case *availda.DAClient:
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
