package main

import (
	"net/http"
	"power4web/modules/modules"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	board := modules.NewBoard()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
