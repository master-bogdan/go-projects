package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Listen   string   `yaml:"listen"`
	Backends []string `yaml:"backends"`
	Balancer string   `yaml:"balancer"`
	Health   struct {
		Path                   string        `yaml:"path"`
		Interval               time.Duration `yaml:"interval"`
		Timeout                int           `yaml:"timeout"`
		PassiveFailuresForOpen int           `yaml:"passiveFailuresForOpen"`
		Cooldown               time.Duration `yaml:"cooldown"`
	}
	Retry struct {
		Max     int `yaml:"max"`
		Backoff int `yaml:"backoff"`
	}
	Timeout struct {
		Read  int `yaml:"read"`
		Write int `yaml:"write"`
		Idle  int `yaml:"idle"`
	}
	Transport struct {
		DialTimeout         int `yaml:"dialTimeout"`
		TLSHandshakeTimeout int `yaml:"tlsHandshakeTimeout"`
		MaxIdlePerHost      int `yaml:"maxIdlePerHost"`
	}
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
