package main

import (
	"fmt"
	"lab-go-chess-engine/models"
)

func main() {
	newGame := model.NewBoard()
	newGame.Print()
	fmt.Println("Chess Engine started!")
}
