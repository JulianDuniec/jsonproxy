package configuration

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

type Server struct {
	Port string `yaml:"port"`
}

type JsonP struct {
	CallbackQueryStringParameterName string `yaml:"callbackQueryStringParameterName"`
}

type Service struct {
	BasePath   string `yaml:"basePath"`
	RemotePath string `yaml:"remotePath"`
}

type Configuration struct {
	Server   Server    `yaml:"server"`
	JsonP    JsonP     `yaml:"jsonp"`
	Services []Service `yaml:"services"`
}

/*
 * Loads the configuration.
 * Panics if it cannot open the file.
 */
func Load(path string) Configuration {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var c Configuration
	err = goyaml.Unmarshal(bytes, &c)
	if err != nil {
		panic(err)
	}
	return c
}
