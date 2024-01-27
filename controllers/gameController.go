package controller

import (
	"bufio"
	"fmt"
	model "lab-go-chess-engine/models"

	// "lab-go-chess-engine/services"
	"os"
	"strings"
)

type GameController struct {
	Board *model.Board
}

func New() *GameController {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Chess Engine started!\nDo you want to start new game? [y/n]: ")
		input, _ := reader.ReadString('\n')

		input = strings.ToLower(strings.TrimSpace(input))

		if input == "y" || input == "yes" {
			board := model.NewBoard()
			fmt.Println("New game started")
			board.Print()
			return &GameController{
				Board: board,
			}
		} else if input == "n" || input == "no" {
			return nil
		}

		fmt.Println("Please enter 'y' or 'n'.")
	}
}

func (gc *GameController) Playing() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter piece coordinate to move (or 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Game ended.")
			break
		}

		moveParts := strings.Fields(input)
		if len(moveParts) != 1 {
			fmt.Println("Invalid input. Please use format 'e2 e4'.")
			continue
		}
	}
}
