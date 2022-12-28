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

func TagExistsInConfig(checkTag string, cfg model.Config) bool {
	for _, typeTag := range cfg.TypeTags {
		if checkTag == typeTag {
			return true
		}
	}

	for _, genreTag := range cfg.GenreTags {
		if checkTag == genreTag {
			return true
		}
	}

	for _, ratingTag := range cfg.Ratings {
		if checkTag == ratingTag {
			return true
		}
	}

	if checkTag == cfg.FavoriteTag {
		return true
	}

	if checkTag == cfg.DidNotFinishTag {
		return true
	}

	return false
}
