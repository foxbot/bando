package bando

import (
	"io/ioutil"
	"log"
	"github.com/BurntSushi/toml"
)

type config struct {
	Host string
	Key string
}

func loadConfig() config {
	var err error
	bytes, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)

	var conf config
	_, err = toml.Decode(data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}