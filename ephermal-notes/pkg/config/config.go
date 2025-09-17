package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Db struct {
		Redis redis.Options
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	setEnv()

	cfg.Server.Host = os.Getenv("SERVER_HOST")
	cfg.Server.Port = os.Getenv("SERVER_PORT")

	return cfg, nil
}

func setEnv() {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(filename))

	envFile, err := os.Open(root)
	if err != nil {
		log.Fatal(err)
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
