package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Status 	string `json:"status"`
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Pelicula con ID: %v => %v", movieID, result)
	message := "Chulito says hello world"
	msg := Message{"bien chuletas", message}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msg)

}

func Ping(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ping")
}