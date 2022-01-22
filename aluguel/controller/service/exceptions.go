package controller

import (
	"encoding/json"
	"net/http"
)

type Exception struct {
	Status   int    `json:"status"`
	Mensagem string `json:"retorno"`
	Err      string `json:"excessao"`
}

func ReponseError(w http.ResponseWriter, status int, mensagem string, err error) {
	var request = Exception{
		Status:   status,
		Mensagem: mensagem,
		Err:      err.Error(),
	}
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(request)
}
