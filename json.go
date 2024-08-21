package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Resposning with error %v : %v", code, msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}
func respondWithJSON(w http.ResponseWriter, code int, payLoad interface{}) {
	data, err := json.Marshal(payLoad)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal json response : %v", payLoad)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
