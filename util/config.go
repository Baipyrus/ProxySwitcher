package util

import (
	"encoding/json"
	"io/fs"
	"path/filepath"

	"io"
	"os"
)

func ReadConfigs(cfgPath string) ([]*Config, error) {
	var configs []*Config

	err := filepath.Walk(cfgPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return nil
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil
		}

		var config *Config
		err = json.Unmarshal(bytes, &config)

		if err != nil {
			return nil
		}

		configs = append(configs, config)
		return nil
	})

	return configs, err
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
