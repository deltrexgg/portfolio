package functions

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(
		template.ParseFiles(
		"templates/index.html",
		"components/desk.html"),
	)

	data := map[string]string{
		"Title": "Hari Nandan Porfolio",
		"PageHeader":  "yooh!",
		"Image": "Desk.PNG",
	}

	tmpl.Execute(w, data)
}