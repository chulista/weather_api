package handlers

import (
	"fmt"
	"net/http"
)

func Options(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Implement Options")
}

func GetLocation(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Implement Options")
}