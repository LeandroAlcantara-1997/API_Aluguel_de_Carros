package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var templates *template.Template
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func LoadTemplates(path string) {
	templates = template.Must(template.ParseGlob(path))
}

func ExecuteTemplate(w http.ResponseWriter, nameFile string, data interface{}) {
	templates.ExecuteTemplate(w, nameFile, data)
}

func GeraCookie(r *http.Request, w http.ResponseWriter, username string) error {

	session, err := store.Get(r, "session")
	if err != nil {
		return fmt.Errorf("Erro ao criar sessao %#v", err)

	}
	session.Values["username"] = username
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("Erro ao salvar cookie %#v", err)
	}
	return nil
}
