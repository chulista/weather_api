// Package weather_api INSERT documentation here
package main

import (
	"fmt"
	"github.com/chulista/weather_api/router"
	"log"
	"net/http"
)

func main(){
	router := router.NewRouter()

	server := http.ListenAndServe(":8080", router)
	fmt.Println("el servidor está corriendo en el http://localhost:8080/hola")
	log.Fatal(server)
}

