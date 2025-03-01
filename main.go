package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func main() {
	apiConfig := &apiConfig{}
	const filtpathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(filtpathRoot))
	handler := http.StripPrefix("/app", fileServer)
	mux.Handle("/app/", apiConfig.middlewareMetricsInc(handler))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)
	mux.HandleFunc("GET /admin/metrics", apiConfig.handlerAdminMetrics)
	mux.HandleFunc("POST /admin/reset", apiConfig.handlerHitsReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filtpathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
