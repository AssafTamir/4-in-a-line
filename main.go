package main

import (
	"fmt"
)

func main() {
	game := Game{[boardSize][boardSize]string{}, ""}
	game.newGame()
	isWinner, winner := game.game(&HumanPlayer{}, &ComputerPlayer{numOfGames}, true)
	if isWinner {
		fmt.Println(winner, " win!")
	} else {
		fmt.Println(gameOver)
	}
}
