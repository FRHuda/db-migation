package yaml

import (
	yamlv2 "gopkg.in/yaml.v2"

	"db-migration/config/model"
)

type Yaml struct {
	model *model.Config
}

func New() *Yaml { return &Yaml{} }

func (y *Yaml) Parse(template []byte) {
	var err error
	err = yamlv2.Unmarshal(template, &y.model)
	if err != nil {
		panic(err)
	}
}

func (y *Yaml) GetScheme(service string, env model.Environtment) *model.Scheme {
	if env == model.EnvProd {
		return y.getScheme(y.model.Production, service)
	}
	if env == model.EnvStaging {
		return y.getScheme(y.model.Staging, service)
	}
	if env == model.EnvLocal {
		return y.getScheme(y.model.Local, service)
	}
	return nil
}

func (y *Yaml) IsEnable(service string, env model.Environtment) bool {
	if env == model.EnvProd {
		return y.isEnable(y.model.Production, service)
	}
	if env == model.EnvStaging {
		return y.isEnable(y.model.Staging, service)
	}
	if env == model.EnvLocal {
		return y.isEnable(y.model.Local, service)
	}
	return false
}

func (y *Yaml) getScheme(services []model.Service, name string) *model.Scheme {
	for _, val := range services {
		if name == val.Name {
			return &val.Scheme
		}
	}

	return nil
}

func (y *Yaml) isEnable(services []model.Service, name string) bool {
	for _, val := range services {
		if name == val.Name {
			return val.Enable
		}
	}

	return false
}
