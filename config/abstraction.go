package config

import "db-migration/config/model"

type Configuration interface {
	Parse(template []byte)
	GetScheme(service string, env model.Environtment) *model.Scheme
	IsEnable(service string, env model.Environtment) bool
}
