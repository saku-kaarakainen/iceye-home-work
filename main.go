package main

import (
	"larvis/internal/config"
)

func main() {
	cfg, err := config.Load("./configs/config.json")

	if err != nil {
		panic(err)
	}

	println(cfg.Symbols[0].Name)
}
