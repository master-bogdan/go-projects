package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Listen   string
	Backends []string
	Balancer string
	Health   struct {
		Path                   string
		Interval               string
		Timeout                int
		PassiveFailuresForOpen int
		Coodown                int
	}
	Retry struct {
		Max     int
		Backoff int
	}
	Timeout struct {
		Read  int
		Write int
		Idle  int
	}
	Transport struct {
		DialTimeout         int
		TLSHandshakeTimeout int
		MaxIdlePerHost      int
	}
}

func New() *Config {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}

	return &config
}
