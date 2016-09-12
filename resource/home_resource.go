package resource

import (
	"net/http"
	"html/template"
)

func Main(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./frontend/index.html")
	template.Execute(w, nil)
}