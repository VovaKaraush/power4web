package game

type Game struct {
	Board  [][]int // Plateau 6x7 || Horizontales X Verticales || lignes X colonnes
	Player int     // Joueur actuel (1 ou 2)
	Winner int     // 0 = pas de gagnant, 1 ou 2 = gagnant
}

var game *Game

func NewGame() *Game { // le tableau est formé en descendant sa generation ressemble a une fleche en diagonale-
	board := make([][]int, 6)
	for i := range board {
		board[i] = make([]int, 7)
	}
	return &Game{Board: board, Player: 1, Winner: 0}
}

func (g *Game) PutPiece(col int, player int) bool { // va renvoyer un bool, vrai pour signaler la reussite de l'operation,
	for i := 5; i >= 0; i-- { // si renvoit false ca veut dire que la collonne choisie etait pleine
		if g.Board[i][col] == 0 {
			g.Board[i][col] = player
			g.Player = 3 - g.Player
			return true
		}
	}
	return false
}

func (g *Game) CheckWin() int {
	rows := len(g.Board)
	cols := len(g.Board[0])

	for lgn := 0; lgn < rows; lgn++ {
		for col := 0; col < cols; col++ {
			player := g.Board[lgn][col]
			if player == 0 {
				continue
			}

			// Horizontal
			if col <= cols-4 &&
				player == g.Board[lgn][col+1] &&
				player == g.Board[lgn][col+2] &&
				player == g.Board[lgn][col+3] {
				return player
			}

			// Vertical
			if lgn <= rows-4 &&
				player == g.Board[lgn+1][col] &&
				player == g.Board[lgn+2][col] &&
				player == g.Board[lgn+3][col] {
				return player
			}

			// Diagonal \
			if lgn <= rows-4 && col <= cols-4 &&
				player == g.Board[lgn+1][col+1] &&
				player == g.Board[lgn+2][col+2] &&
				player == g.Board[lgn+3][col+3] {
				return player
			}

			// Diagonal /
			if lgn >= 3 && col <= cols-4 &&
				player == g.Board[lgn-1][col+1] &&
				player == g.Board[lgn-2][col+2] &&
				player == g.Board[lgn-3][col+3] {
				return player
			}
		}
	}
	return 0
}
