package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

func ReadFile(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var cfg Config
	return &cfg, yaml.Unmarshal(data, &cfg)
}
