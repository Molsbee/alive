package resource

import (
	"net/http"
	"html/template"
)

var t, _ = template.ParseFiles(
	"./frontend/templates/index.html",
	"./frontend/templates/monitor-create.html",
	"./frontend/templates/navbar.html",
	"./frontend/templates/sidebar.html",
)

func Main(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.html", nil)
}

func CreateMonitor(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "monitor-create.html", nil)
}