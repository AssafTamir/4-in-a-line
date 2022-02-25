package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type simulationResult struct {
	countWins int
	index     int
}
type Player interface {
	move(game *Game)
}

type ComputerPlayer struct {
	numOfGames int
}

type RandomPlayer struct {
	random *rand.Rand
}

type HumanPlayer struct{}

func (player *ComputerPlayer) move(game *Game) {
	start := time.Now()
	ch := make(chan simulationResult, boardSize)
	var wg sync.WaitGroup

	maxWins := -1
	maxWinsIndex := -1
	wg.Add(boardSize)
	for i := 0; i < boardSize; i++ {
		go func(i int) {
			defer wg.Done()
			if game.isPossibleMove(i) {
				countWins := 0
				var s = rand.NewSource(time.Now().UnixNano())
				var random = rand.New(s)
				for games := 0; games < player.numOfGames; games++ {
					newGame := *game
					newGame.applyMove(i).switchTurn()
					isWinner, winner := newGame.game(&RandomPlayer{random}, &RandomPlayer{random}, false)
					if isWinner && winner == O {
						countWins++
					}
				}
				ch <- simulationResult{countWins, i}
			}
		}(i)
	}
	wg.Wait()
	close(ch)
	for wins := range ch {
		if wins.countWins > maxWins {
			maxWins = wins.countWins
			maxWinsIndex = wins.index
		}

	}
	game.applyMove(maxWinsIndex)
	elapsed := time.Since(start)
	fmt.Printf("Turn took %s", elapsed)
}

func (player *HumanPlayer) move(game *Game) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(game.turn + " Enter move: ")
		text, _ := reader.ReadString('\n')
		move, _ := strconv.Atoi(text[0:1])
		if game.applyMove(move-1) != nil {
			return
		}
	}
}
func (player *RandomPlayer) move(game *Game) {

	for {
		move := player.random.Intn(boardSize)
		if game.isPossibleMove(move) {
			if game.applyMove(move) != nil {
				return
			}
		}

	}
}
