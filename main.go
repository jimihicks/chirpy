package main

import (
	"log"
	"net/http"
)

func main() {
	const filtpathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(filtpathRoot)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filtpathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
