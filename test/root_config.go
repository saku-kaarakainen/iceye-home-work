package main

import (
	"larvis/internal/card"
	"larvis/internal/game"
	"larvis/internal/hand"
)

// Copy of the struct from main.go
// for the integration tests

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
