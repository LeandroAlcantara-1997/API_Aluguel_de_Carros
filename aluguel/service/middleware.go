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
			ReponseError(w, 401, "Token inválido", fmt.Errorf("Token não pode estar vazio"))
			return
		}

		token, err := extraiToken(r.Header.Get("Authorization"))
		if err != nil {
			ReponseError(w, http.StatusUnauthorized, "Token inválido", err)
		}
		
		if err := ValidaCliente(token); err != nil {
			ReponseError(w, 401, "token inválido", err)
			return
		}
		next(w, r)
	}

}


func MiddlewareAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contain := strings.Contains(r.Header.Get("Authorization"), bearer_schema); !contain {
			ReponseError(w, 401, "Token inválido", fmt.Errorf("Token não pode estar vazio"))
			return
		}

		token, err := extraiToken(r.Header.Get("Authorization"))
		if err != nil {
			ReponseError(w, 401, "Token inválido", err)
		}
		
		if err := ValidaAdmin(token); err != nil {
			ReponseError(w, 401, "token inválido", err)
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