package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type EnvVariables struct {
	HOST string `env:"HOST,required"`
}

func LoadEnv(paths ...string) *EnvVariables {
	path := ".env"
	if len(paths) > 0 {
		path = paths[0]
	}

	if err := godotenv.Load(path); err != nil {
		panic(fmt.Errorf("error loading %s file", path))
	}

	var config EnvVariables
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		panic(err)
	}

	return &config
}
