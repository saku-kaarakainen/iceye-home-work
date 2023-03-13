package main

import (
	"fmt"
	"time"
)

type Event struct {
	Type string
	Data interface{}
}

func main() {
	fmt.Println("Welcome to poker game!")

	cfg, err := LoadConfig("./configs/config.json")
	if err != nil {
		panic(err)
	}

	color := cfg.Domains.Card.Colors[0].Name
	fmt.Println("Color:")
	fmt.Println(color)

	eventChan := make(chan Event)

	// start game loop in a separate goroutine
	go gameLoop(eventChan)

	// wait for game loop to end
	<-eventChan
	fmt.Println("Thank you for playing with us!")
}

func gameLoop(eventChan chan Event) {
	for {
		// process events
		select {
		case event := <-eventChan:
			handleEvent(event)
		default:
			// no events, do game update
			update()
		}

		// sleep for some time before next update
		time.Sleep(100 * time.Millisecond)
	}
}

func handleEvent(event Event) {
	fmt.Println("Handling event:", event.Type)

	switch event.Type {
	case "player_joined":
		fmt.Printf("Player %s joined the game\n", event.Data)
	case "player_left":
		fmt.Printf("Player %s left the game\n", event.Data)
	case "bet_placed":
		fmt.Printf("Player %s placed a bet of %d chips\n", event.Data.(string), 100)
	}
}

func update() {
	// do game update
	fmt.Println("Updating game...")
}
