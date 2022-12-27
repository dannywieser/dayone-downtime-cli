package config

import (
	"dod/pkg/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadConfig() (*model.Config, error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config model.Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
