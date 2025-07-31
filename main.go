package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mourasjames/cep-temp-api/handler"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/weather", handler.WeatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Servidor iniciado em :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
