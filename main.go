package main

import (
	"net/http"
	"text/template"
)

type Game struct {
	Board         [6][7]int // 0=vide, 1=rouge, 2=jaune
	CurrentPlayer int
}

type GamePageData struct {
	PageTitle     string
	Board         [6][7]int
	CurrentPlayer int
	Winner        int
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Board := 1
		data := GamePageData{
			PageTitle:     "Puissance 4",
			Board:         Game.Board,
			CurrentPlayer: Game.CurrentPlayer,
			Winner:        CheckWinner(),
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
