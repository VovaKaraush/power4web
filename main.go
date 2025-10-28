package main

import "fmt"

type Game struct {
	Board  [][]int // Plateau 6x7 || Horizontales X Verticales || lignes X colonnes
	Player int     // Joueur actuel (1 ou 2)
	Winner int     // 0 = pas de gagnant, 1 ou 2 = gagnant
}

func NewGame() *Game { // le tableau est formé en descendant sa generation ressemble a une fleche en diagonale-
	board := make([][]int, 6)
	for i := range board {
		board[i] = make([]int, 7)
	}
	return &Game{Board: board, Player: 1, Winner: 0}
}

func (g *Game) PutPiece(col int, player int) bool { // va renvoyer un bool, vrai pour signaler la reussite de l'operation,
	for i := 5; i > 0; i-- { // si renvoit false ca veut dire que la collonne choisie etait pleine
		if g.Board[i][col] == 0 {
			return true
		}
	}
	return false
}

func (g *Game) checkWin() int { //horizontal then vertical then diagonal+ and then diagonal-
	for col := 0; col < 6; col++ { // verticales
		for lgn := 5; lgn >= 0; lgn-- { // horizontales
			if g.Board[lgn][col] != 0 {
				checked := g.Board[lgn][col]
				if g.Board[lgn][col] == g.Board[lgn][col+1] && g.Board[lgn][col+1] == g.Board[lgn][col+2] && g.Board[lgn][col+2] == g.Board[lgn][col+3] { // check de win horizontal
					return (checked)
				} else if g.Board[lgn][col] == g.Board[lgn+1][col] && g.Board[lgn+1][col] == g.Board[lgn+2][col] && g.Board[lgn+2][col] == g.Board[lgn+3][col] { // check horizontal
					return (checked)
				} else if g.Board[lgn][col] == g.Board[lgn+1][col+1] && g.Board[lgn+1][col+1] == g.Board[lgn+2][col] && g.Board[lgn+2][col+2] == g.Board[lgn+3][col+3] { // check diagonal +
					return (checked)
				} else if g.Board[lgn][col] == g.Board[lgn-1][col+1] && g.Board[lgn-1][col+1] == g.Board[lgn-2][col] && g.Board[lgn-2][col+2] == g.Board[lgn-3][col+3] { // check diagonal -
					return (checked)
				}
			}
		}
		return 0
	}
}

func InputScan() int {
	var x int
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Entrée invalide, réessaie.")
		AssignValue()
	}
	return x
}

func main() {
	var game *Game = NewGame()
	for game.Winner == 0 {
		fmt.Print("Joueur ")
		fmt.Print(game.Player)
		fmt.Print(" choisis la colonne dans laquelle tu veux mettre la piece")
		x := InputScan()
		game.PutPiece(x, game.Player)
	}
}
