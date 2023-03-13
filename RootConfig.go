package main

type RootConfig struct {
	DeckConfigs DeckConfig `json:"deckConfig"`       
	GameConfigs GameConfig `json:"gameConfig"`
}
