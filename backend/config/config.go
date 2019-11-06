package config

import (
	"github.com/go-playground/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Conf *Config

func Setup() {
	data, err := read("./config.yaml")
	if err != nil {
		log.WithError(err).
			Fatal("Couldn't read configures")
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		log.WithError(err).
			Fatal("Couldn't unmarshal configures")
	}

}

func read(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
