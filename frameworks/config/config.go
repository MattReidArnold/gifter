package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/mattreidarnold/gifter/domain/errgroup"
)

type Config struct {
	MongoDatabase string
	MongoHost     string
	MongoPassword string
	MongoPort     string
	MongoUsername string
}

var ErrEnvNotFound error = errors.New("required env var not found")

func Init() (*Config, error) {
	errGroup := errgroup.NewErrorGroup("ConfigErrors")

	mongoDatabase := requireEnv("MONGO_DATABASE", errGroup)
	mongoHost := requireEnv("MONGO_HOST", errGroup)
	mongoPassword := requireEnv("MONGO_PASSWORD", errGroup)
	mongoPort := requireEnv("MONGO_PORT", errGroup)
	mongoUsername := requireEnv("MONGO_USERNAME", errGroup)

	if !errGroup.Empty() {
		return &Config{}, errGroup
	}

	return &Config{
		MongoDatabase: mongoDatabase,
		MongoHost:     mongoHost,
		MongoPassword: mongoPassword,
		MongoPort:     mongoPort,
		MongoUsername: mongoUsername,
	}, nil
}

func requireEnv(key string, errGroup errgroup.ErrorGroup) string {
	value, found := os.LookupEnv(key)
	if !found {
		errGroup.Append(fmt.Errorf("%s: %w", key, ErrEnvNotFound))
		return ""
	}
	return value
}
