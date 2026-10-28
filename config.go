package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type backendConfig struct {
	Address string `yaml:"address"`
}

type config struct {
	Address  string          `yaml:"address"`
	Mode     string          `yaml:"mode"`
	Backends []backendConfig `yaml:"backends"`
}

func loadConfig(path string) (*config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
