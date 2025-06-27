package util

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"io"
	"os"
)

func (c *Config) Delete() error {
	return os.Remove(c.filepath)
}

func ReadConfigs(cfgPath string, examples bool) ([]*Config, error) {
	var configs []*Config

	err := filepath.Walk(cfgPath, func(path string, info fs.FileInfo, err error) error {
		name := info.Name()
		isExample := strings.HasSuffix(name, ".example.json")
		notJson := !strings.HasSuffix(name, ".json")
		if info.IsDir() || (isExample && !examples) || notJson {
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

		config.filepath = path
		configs = append(configs, config)
		return nil
	})

	return configs, err
}

func SaveConfig(cfgPath string, config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// 'WSL - Sudo - NPM' turns into 'wsl_sudo_npm'
	name :=
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ToLower(config.Name),
				" ", ""),
			"-", "_")

	cfgName := fmt.Sprintf("%s.json", name)
	cfgFile := filepath.Join(cfgPath, cfgName)
	err = os.WriteFile(cfgFile, data, 0666)

	return err
}
