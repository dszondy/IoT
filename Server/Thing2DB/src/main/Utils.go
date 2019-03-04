package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonError struct {
	Error string `json:"error"`
	Success bool `json:"success"`
}

type jsonSuccess struct {
	Success bool `json:"success"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, jsonError{Error: message, Success: false})
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}
func respondOK(w http.ResponseWriter){
	respondWithJSON(w, http.StatusOK, jsonSuccess{true})
}

func tryReadRequestBody(r *http.Request, payload interface{}, w http.ResponseWriter) bool {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(payload); err != nil {
		log.Printf("Invalid request payload: %v", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	defer r.Body.Close()
	return true
}
