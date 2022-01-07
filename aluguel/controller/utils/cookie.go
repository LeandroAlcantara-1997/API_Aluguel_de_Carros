package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

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

func GetSecao(r *http.Request) error {
	// grab the session
	session, err := store.Get(r, "session")
	if err != nil {
		return fmt.Errorf("Erro ao pegar seção, ", err)
	}
	// grab the username from the session object
	untyped, ok := session.Values["username"]
	// Verify that the username was actually there
	if !ok {
		return nil
	}
	// quick type assertion because the object store the data as
	// empty interface
	_, ok = untyped.(string)
	if !ok {
		return nil
	}
	return nil
}
