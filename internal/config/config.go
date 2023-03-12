package config

import (
	"encoding/json"
	"io/ioutil"
)


func Load(filename string) (Config, error) {
	jsonBlob, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	json.Unmarshal(jsonBlob, &config)

	return config, nil
}
