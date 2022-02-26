package main

import "fmt"

func main() {
	player1 := HumanPlayer{X}
	player2 := ComputerPlayer{1, O}
	player3 := ComputerPlayer{numOfGames, W}

	game := Game{[boardSize][boardSize]*Token{}, []Player{&player1, &player2, &player3}, 0}
	game.newGame()
	isWinner, winner := game.game(true)
	if isWinner {
		fmt.Println(winner.char, " win!")
	} else {
		fmt.Println(gameOver)
	}
}

//func main() {
//	player1 := HumanPlayer{X}
//	player2 := ComputerPlayer{numOfGames, O}
//	game := Game{[boardSize][boardSize]*Token{}, []Player{&player1, &player2}, 0, nil}
//	game.newGame()
//	isWinner, winner := game.game(true)
//	if isWinner {
//		fmt.Println(winner.char, " win!")
//	} else {
//		fmt.Println(gameOver)
//	}
//}
