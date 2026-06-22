package main

import (
	"net/http"

	"github.com/deltrexgg/profolio/functions"
)

func main() {

	http.HandleFunc("/", functions.HomePage)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	println("Server running :8080")

	http.ListenAndServe(":8080", nil)
}
