package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"github.com/mamemomonga/go-study-simple-toot/src/don"
)

// Config Config
type Config struct {
	ClientName string        `yaml:"client_name"`
	Mastodon   don.UserLogin `yaml:"mastodon"`
}

func configLoad(filename string) (cnf Config, err error) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return cnf, err
	}

	err = yaml.Unmarshal(buf, &cnf)
	if err != nil {
		return cnf, err
	}

	log.Printf("trace: Load %s", filename)

	return cnf, nil
}
