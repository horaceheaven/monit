package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()

	port := ":8080"

	envPort := os.Getenv("MONIT_PORT")

	if envPort != "" {
		port = ":" + envPort
		log.Print("using port from environment variable ", port)
	} else {
		log.Print("using default port ", port)
	}

	log.Fatal(http.ListenAndServe(port, router))
}
