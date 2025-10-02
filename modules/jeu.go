package modules

const (
	rows    = 6
	columns = 7
)

func newBoard() Board {
	var b Board
	for i := range b {
		for j := range b[i] {
			b[i][j] = 0
		}
	}
	return b
}

// joueurs 1 et 2. 1 sera bleu, 2 sera jaune.
// on commencera le jeu avec l'action du joueur 1 et nous allons suivre
// l'avancement du jeu avec modulo du tour actuel. si tour modulo 2 rends
// qqch cest au tour du joueur 1, si ca rend 0 c'est au tour du joueur 2.

func whoseMove()
