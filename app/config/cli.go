package config

import "github.com/caarlos0/env"

type CLI struct {
	GRPCServerAddress string `env:"GRPC_SERVER_ADDRESS,required"`
}

func NewCLI() (*CLI, error) {
	var cfg CLI

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
