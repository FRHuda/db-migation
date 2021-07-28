package config

import "db-migration/config/model"

type Configuration interface {
	Parse(template []byte)
	GetScheme(service, env string) *model.Scheme
}
