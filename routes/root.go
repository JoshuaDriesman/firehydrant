package routes

import (
	"html/template"
	"net/http"
)

type RootPageData struct {
	Version string
}

func Root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data := RootPageData{
		Version: "0.0.1",
	}
	tmpl.Execute(w, data)
}
