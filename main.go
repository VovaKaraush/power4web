package main

import (
    "net/http"
    "text/template"
)

type Game struct {
    Board         [6][7]int // 0 = vide, 1 = rouge, 2 = jaune
    CurrentPlayer int
}

type GamePageData struct {
    PageTitle     string
    Board         [6][7]int
    CurrentPlayer int
    Winner        int
}


func CheckWinner(board [6][7]int) int {
    return 0 // 0 = pas de gagnant
}

func main() {
    game := &Game{}
    tmpl := template.Must(template.ParseFiles("layout.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := GamePageData{
            PageTitle:     "Puissance 4",
            Board:         game.Board,
            CurrentPlayer: game.CurrentPlayer,
            Winner:        CheckWinner(game.Board),
        }
        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":8080", nil)
}
