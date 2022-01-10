package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Addr string `json:"address"`
	Port string `json:"port"`
}

func getAddrFromConfig(file string) string {
	cfg := &config{}

	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}

	err = json.Unmarshal(jsonFile, cfg)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port)
}