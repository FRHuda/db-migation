package config

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/sethvargo/go-envconfig"

	"db-migration/config/model"
	"db-migration/config/yaml"
)

type Config struct {
	DbMigration *DbMigration
}
type DbMigration struct {
	DSNMigrationLocal      string `env:"POST_MIGRATE_DSN_STAGING,default=host=localhost port=5432 database=testdb user=postgres password=postgres sslmode=disable"`
	DSNMigrationStaging    string `env:"POST_MIGRATE_DSN_STAGING,default=host=localhost port=5432 database=testdb user=postgres password=postgres sslmode=disable"`
	DSNMigrationProduction string `env:"POST_MIGRATE_DSN_PRODUCTION,default=postgres://postgres@localhost:5432/sellfazz?sslmode=disable"`
	SourceMigrationName    string `env:"POST_DB_SOURCE_MIGRATION_NAME,default=go.rice"`
}

func Process(c *Config) error {
	return envconfig.Process(context.Background(), c)
}

func GetScheme(service string, env string, configType string, configLocation string) *model.Scheme {
	var err error

	if configType == "" {
		configType = model.TypeYaml
	}

	fmt.Println("Reading Config ", configLocation)

	byteFile, err := ioutil.ReadFile(configLocation)
	if err != nil {
		panic(err)
	}

	parser := getImplementation(configType)

	if parser == nil {
		panic("cannot_find_parser")
	}

	parser.Parse(byteFile)

	if env == "" {
		panic("ENV is NULL")
	}
	return parser.GetScheme(service, env)
}

func getImplementation(configType string) Configuration {
	if configType == model.TypeYaml {
		return yaml.New()
	}

	return nil
}
