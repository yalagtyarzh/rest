package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents the config of the launched http server
type Config struct {
	IP   string `env:"REST_IP" env-default:"127.0.0.1"`
	Port string `env:"REST_PORT" env-default:":8080"`
}

var cfg *Config
var once sync.Once

// GetConfig gets the config from the environment variables and writes the values to the Config structure, returning it
func GetConfig() *Config {
	once.Do(
		func() {
			cfg = &Config{}

			if err := cleanenv.ReadEnv(cfg); err != nil {
				fmt.Printf("environment is not OK: %s\n", err)
				os.Exit(1)
			}
		},
	)

	return cfg
}
