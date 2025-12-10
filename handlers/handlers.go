package handlers

import (
	"html/template"
	"net/http"
	"power4web/game"
	"strconv"
)

var currentGame *game.Game

// SetGame initializes the handlers with a game instance
func SetGame(g *game.Game) {
	currentGame = g
}

// HomeHandler serves the game page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("game.html"))
	tmpl.Execute(w, currentGame)
}

// PlayHandler handles game moves
func PlayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && currentGame.Winner == 0 {
		colStr := r.FormValue("column")
		col, _ := strconv.Atoi(colStr)

		if col >= 0 && col < 7 {
			if currentGame.PutPiece(col, currentGame.Player) {
				currentGame.Winner = currentGame.CheckWin()
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ResetHandler resets the game
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
