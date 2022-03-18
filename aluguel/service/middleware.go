package service

import (
	"net/http"
	"strings"
)

func MiddlewareCliente(http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const Bearer_schema = "Bearer "
		header := r.Header.Get("Authorization")
		if header == "" {
			ReponseError(w, 400, "token inválido", nil)
			return
		}

		token := strings.Split(header, Bearer_schema)
		if !ValidateToken(token[1]) {
			ReponseError(w, 400, "token inválido", nil)
			return
		}
	}
}
