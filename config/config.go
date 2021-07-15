package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Server struct {
	Http Http `json:"http" yaml:"http"`
	Mysql Mysql `json:"mysql" yaml:"mysql"`
}

//read yaml config
func GetConfig() (*Server, error){
	conf := &Server{}
	if f, err := os.Open("config.yaml"); err != nil {
		return nil, err
	} else {
		_ = yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}
