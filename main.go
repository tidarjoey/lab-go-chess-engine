package main

import (
	// "lab-go-chess-engine/models"
	// "lab-go-chess-engine/services"
	"fmt"
	controller "lab-go-chess-engine/controllers"
)

func main() {
	game := controller.New()
	if game == nil {
		fmt.Println("Game ended.")
	}
	game.Playing()
}
