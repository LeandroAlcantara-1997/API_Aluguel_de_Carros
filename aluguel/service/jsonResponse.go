package service

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, obj interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}