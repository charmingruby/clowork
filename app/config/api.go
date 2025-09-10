package config

import "github.com/caarlos0/env"

type API struct {
	RestServerPort    string `env:"REST_SERVER_PORT,required"`
	GRPCServerAddress string `env:"GRPC_SERVER_ADDRESS,required"`
	PostgresURL       string `env:"POSTGRES_URL,required"`
	LogLevel          string `env:"LOG_LEVEL"`
	DatabasePageSize  int    `env:"DATABASE_PAGE_SIZE,required"`
}

func NewAPI() (*API, error) {
	var cfg API

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
