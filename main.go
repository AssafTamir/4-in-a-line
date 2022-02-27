package main

import "fmt"

func main() {
	players := []Player{
		&ComputerPlayer{1, X},
		&ComputerPlayer{1, O},
		&ComputerPlayer{1, W},
		&ComputerPlayer{10000, D},
	}
	counter := make(map[string]int)
	for _, p := range players {
		counter[p.getToken().char] = 0
	}
	for i := 0; i < 100; i++ {
		game := Game{[boardSize][boardSize]*Token{}, players, 0}
		game.newGame()
		isWinner, winner := game.game(false)
		if isWinner {
			fmt.Println(winner.char, " win!")
			counter[winner.char]++
		} else {
			fmt.Println(gameOver)
		}
	}
	fmt.Println(counter)
}
