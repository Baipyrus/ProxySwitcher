package util

import (
	"encoding/json"
	"io"
	"os"
)

func ReadConfigs() ([]Config, error) {
	file, readErr := os.Open("configs.json")

	if readErr != nil {
		return nil, readErr
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var config []Config
	unmarshalErr := json.Unmarshal(bytes, &config)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return config, nil
}
