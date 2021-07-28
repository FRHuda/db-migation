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

func (y *Yaml) GetScheme(service, env string) *model.Scheme {
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

func (y *Yaml) getScheme(services []model.Service, name string) *model.Scheme {
	for _, val := range services {
		if name == val.Name {
			return &val.Scheme
		}
	}

	return nil
}
