package main

import (
	"fmt"
	"net/http"
	"power4web/game"
	"power4web/handlers"
)

func main() {
	// Initialize game and pass it to handlers
	gameInstance := game.NewGame()
	handlers.SetGame(gameInstance)

	// Register HTTP routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/play", handlers.PlayHandler)
	http.HandleFunc("/reset", handlers.ResetHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
