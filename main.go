package main

import (
	"net/http"

	"github.com/deltrexgg/profolio/functions"
)

func main() {
	port := "8090"

	http.HandleFunc("/", functions.HomePage)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	println("Server running :",port)

	http.ListenAndServe(":"+port, nil)
}
