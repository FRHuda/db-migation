package model

const (
	EnvProd    = "production"
	EnvStaging = "staging"
	EnvLocal   = "local"
	TypeYaml   = "yaml"
	TypeJson   = "json"
)

type Config struct {
	Production []Service `json:"production" yaml:"production"`
	Staging    []Service `json:"staging" yaml:"staging"`
	Local      []Service `json:"local" yaml:"local"`
}

type Service struct {
	Name   string `yaml:"name" json:"name"`
	Scheme Scheme `yaml:"scheme" json:"scheme"`
}

type Scheme struct {
	Migration int `yaml:"migration" json:"migration"`
}
