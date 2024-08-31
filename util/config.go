package util

import (
	"encoding/json"
	"io"
	"os"
)

func ReadConfigs() ([]*Config, error) {
	file, err := os.Open("configs.json")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var config []*Config
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}
	return config, nil
}
