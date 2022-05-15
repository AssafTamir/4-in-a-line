package main

import "fmt"

func main() {
	players := []Player{
		&HumanPlayer{X},
		&ComputerPlayer{1, O},
	}
	counter := make(map[string]int)
	for _, p := range players {
		counter[p.getToken().char] = 0
	}
	for i := 0; i < 100; i++ {
		game := Game{[boardSize][boardSize]*Token{}, players, 0, 0}
		game.newGame()
		isWinner, winner := game.game(true)
		if isWinner {
			fmt.Println(winner.char, " win!, Number of moves: ", game.moveCounter)
			counter[winner.char]++

		} else {
			fmt.Println(gameOver)
		}
	}
	fmt.Println(counter)
}
