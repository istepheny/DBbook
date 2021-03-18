package config

import (
	"dbbook/pkg/helper"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Ticker    int        `json:"ticker"`
	Server    Server     `json:"server"`
	Databases []Database `json:"databases"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Load() (config Config) {
	fh, err := os.Open(helper.ConfigFilePath())
	defer fh.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Fatal(err)
		return
	}

	return config
}
