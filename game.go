package main

import "fmt"

type Game struct {
	board   [boardSize][boardSize]*Token
	player1 Player
	player2 Player
	turn    Player
}

func (game *Game) newGame() *Game {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			game.board[i][j] = OpenPosition
		}
	}
	game.turn = game.player1
	return game
}

func (game *Game) isBoardFull() bool {
	for i := 0; i < boardSize; i++ {
		if game.board[0][i] == OpenPosition {
			return false
		}
	}
	return true
}
func (game *Game) applyMove(move int) *Game {
	if !game.isPossibleMove(move) {
		return nil
	}
	for i := 7; i >= 0; i-- {
		if game.board[i][move] == OpenPosition {
			game.board[i][move] = game.turn.getToken()
			return game
		}

	}
	return nil
}
func (game *Game) isPossibleMove(pos int) bool {
	return game.board[0][pos] == OpenPosition
}

func (game *Game) String() (ret string) {

	ret += string(colorPurple) + "\n|  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |" + string(colorReset) + "\n"
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			ret += "|  "
			ret += game.board[i][j].string()
			ret += "  "
		}
		ret += "|\n"
	}
	return ret
}
func (game *Game) isWinnerOnPosition(x int, y int) bool {
	for i := 1; i < 4; i++ {
		if x+i >= boardSize || game.board[x][y] != game.board[x+i][y] {
			break
		}
		if i == 3 {
			return true
		}
	}
	for i := 1; i < 4; i++ {
		if y+i >= boardSize || game.board[x][y] != game.board[x][y+i] {
			break
		}
		if i == 3 {
			return true
		}
	}
	for i := 1; i < 4; i++ {
		if x+i >= boardSize || y+i >= boardSize || game.board[x][y] != game.board[x+i][y+i] {
			break
		}
		if i == 3 {
			return true
		}
	}

	for i := 1; i < 4; i++ {
		if x-i < 0 || y+i >= boardSize || game.board[x][y] != game.board[x-i][y+i] {
			break
		}
		if i == 3 {
			return true
		}
	}

	return false
}
func (game *Game) isGameOver() (bool, *Token) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if game.board[i][j] == OpenPosition {
				continue
			} else {
				if game.isWinnerOnPosition(i, j) {
					return true, game.board[i][j]
				}
			}
		}
	}
	return false, nil
}
func (game *Game) game(isPrintBoard bool) (bool, *Token) {

	for {
		if isPrintBoard {
			fmt.Println(game)
		}

		isWinner, token := game.isGameOver()
		if isWinner {
			return true, token
		}
		if game.isBoardFull() {
			return false, nil
		}
		game.turn.move(game)
		game.switchTurn()
	}
}

func (game *Game) switchTurn() *Game {
	if game.turn.getToken() == X {
		game.turn = game.player2
	} else {
		game.turn = game.player1
	}
	return game
}
