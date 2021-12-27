package controller

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates(path string) {
	templates = template.Must(template.ParseGlob(path))
}

func ExecuteTemplate(w http.ResponseWriter, nameFile string, data interface{}) {
	templates.ExecuteTemplate(w, nameFile, data)
}
