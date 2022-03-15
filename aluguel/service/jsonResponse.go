package service

import (
	"encoding/json"
	"net/http"
	"time"
)

type exception struct {
	Status    int       `json:"status"`
	Mensagem  string    `json:"mensagem"`
	Err       string    `json:"excessao"`
	Timestamp time.Time `json:"timestamp"`
}

func ReponseError(w http.ResponseWriter, status int, mensagem string, err error) {
	w.Header().Set("Content-Type", "application/json")
	var request = exception{
		Status:    status,
		Mensagem:  mensagem,
		Err:       err.Error(),
		Timestamp: time.Now(),
	}
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(request)
}

func JsonResponse(w http.ResponseWriter, status int, obj interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}