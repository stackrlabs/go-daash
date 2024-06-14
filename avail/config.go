package avail

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Seed               string `json:"seed"`
	WsRpcURL           string `json:"wsRpcUrl"`
	HttpApiURL         string `json:"httpApiUrl"`
	AppID              int    `json:"app_id"`
	DestinationDomain  int    `json:"destination_domain"`
	DestinationAddress string `json:"destination_address"`
	Timeout            int    `json:"timeout"`
	Network            string `json:"network"`
}

func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return fmt.Errorf("cannot open config file:%w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("cannot read config file:%w", err)
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file:%w", err)
	}

	return nil
}
