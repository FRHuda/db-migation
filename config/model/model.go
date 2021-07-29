package model

type Environtment string

const (
	EnvProd    Environtment = "production"
	EnvStaging Environtment = "staging"
	EnvLocal   Environtment = "local"
	TypeYaml                = "yaml"
	TypeJson                = "json"
)

var Environtments = []Environtment{
	EnvLocal,
	EnvStaging,
	EnvProd,
}

type Config struct {
	Production []Service `json:"production" yaml:"production"`
	Staging    []Service `json:"staging" yaml:"staging"`
	Local      []Service `json:"local" yaml:"local"`
}

type Service struct {
	Name   string `yaml:"name" json:"name"`
	Enable bool   `yaml:"enable" json:"enable"`
	Scheme Scheme `yaml:"scheme" json:"scheme"`
}

type Scheme struct {
	Migration int `yaml:"migration" json:"migration"`
}
