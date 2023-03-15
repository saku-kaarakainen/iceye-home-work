package main

import (
	"encoding/json"
	"io/ioutil"
	"larvis/internal/card"
	"larvis/internal/hand"
	"larvis/internal/game"
)

type DomainConfig struct {
	// game
	Game game.GameConfig `json:"game"`

	// hand
	Hand hand.HandConfig `json:"hand"`

	// card
	Card card.CardConfig `json:"card"`
}

type RootConfig struct {
	Domains DomainConfig `json:"domains"`
}

// Load reads the configuration file with the given filename and returns a RootConfig.
// If the file is not found or cannot be read, Load panics.
func LoadConfig(filename string) (RootConfig, error) {
	jsonBlob, err := ioutil.ReadFile(filename)
	if err != nil {
		return RootConfig{}, err
	}

	var config RootConfig
	json.Unmarshal(jsonBlob, &config)

	return config, nil
}
