package service

import (
	"fmt"
	"net/http"
	"strings"
)

const bearer_schema = "Bearer "

func MiddlewareCliente(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contain := strings.Contains(r.Header.Get("Authorization"), bearer_schema); !contain {
			ReponseError(w, 400, "Token inválido", fmt.Errorf("Token não pode estar vazio"))
			return
		}

		token, err := extraiToken(r.Header.Get("Authorization"))
		if err != nil {
			ReponseError(w, http.StatusUnauthorized, "Token inválido", err)
		}
		
		if err := ValidateToken(token); err != nil {
			ReponseError(w, 400, "token inválido", err)
			return
		}
		next(w, r)
	}

}

func extraiToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("token não pode estar vazio")
	}
	
	token := strings.Split(header, bearer_schema)
	return token[1], nil
}