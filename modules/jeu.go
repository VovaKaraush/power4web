package modules

type Game struct {
	Board [6][7]int // 0=vide, 1=rouge, 2=jaune
	tour  int
	// 7 colonnes et 6 lignes
}

func (g Game) addChip(clmn int, p int) { // clmn = colonne
	
}

func Play(b Game) {
	turn := 1
	for {
		if turn%2 != 0 {

		}
	}
}

func NewBoard() Game.Board {
	var b Game.Board
	for i := range b {
		for j := range b[i] {
			b[i][j] = 0
		}
	}
	return b
}

func addChip(b Board) {

}

func takeTurn(p) { //action de mettre une piece. p pour player(definit le jouer qui fait le coup)

}

// joueurs 1 et 2. 1 sera bleu, 2 sera jaune.
// on commencera le jeu avec l'action du joueur 1 et nous allons suivre
// l'avancement du jeu avec modulo du tour actuel. si tour modulo 2 rends
// qqch cest au tour du joueur 1, si ca rend 0 c'est au tour du joueur 2.
