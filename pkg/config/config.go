package config

import (
	"dbbook/pkg/helper"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Load() (databases []Database) {
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

	err = json.Unmarshal(jsonData, &databases)
	if err != nil {
		log.Fatal(err)
		return
	}

	return databases
}
