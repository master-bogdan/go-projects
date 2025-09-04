package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Listen   string   `yaml:"listen"`
	Backends []string `yaml:"backends"`
	Balancer string   `yaml:"balancer"`
	Health   struct {
		Path                   string `yaml:"path"`
		Interval               int    `yaml:"interval"`
		Timeout                int    `yaml:"timeout"`
		PassiveFailuresForOpen int    `yaml:"passiveFailuresForOpen"`
		Cooldown               int    `yaml:"cooldown"`
	}
	Retry struct {
		Max     int `yaml:"max"`
		Backoff int `yaml:"backoff"`
	} `yaml:"retry"`
	Timeout struct {
		Read  int `yaml:"read"`
		Write int `yaml:"write"`
		Idle  int `yaml:"idle"`
	} `yaml:"timeout"`
	Transport struct {
		DialTimeout         int `yaml:"dialTimeout"`
		TLSHandshakeTimeout int `yaml:"tlsHandshakeTimeout"`
		MaxIdlePerHost      int `yaml:"maxIdlePerHost"`
	} `yaml:"transport"`
}

func New() *Config {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
