package util

import (
	"encoding/json"
	"io"
	"os"
)

func ReadConfigs(name string) ([]*Config, error) {
	file, err := os.Open(name)

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

func SaveConfig(name string, config Config) error {
	configs, _ := ReadConfigs(name)
	configs = append(configs, &config)

	data, err := json.Marshal(configs)
	if err != nil {
		return err
	}

	err = os.WriteFile(name, data, 0666)
	return err
}
