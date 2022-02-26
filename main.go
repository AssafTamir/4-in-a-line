package main

import (
	"fmt"
)

func main() {
	player1 := HumanPlayer{X}
	player2 := ComputerPlayer{numOfGames, O}

	game := Game{[boardSize][boardSize]*Token{}, &player1, &player2, &player1}
	game.newGame()
	isWinner, winner := game.game(true)
	if isWinner {
		fmt.Println(winner.char, " win!")
	} else {
		fmt.Println(gameOver)
	}
}
