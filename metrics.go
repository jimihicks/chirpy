package main

import (
	"fmt"
	"io"
	"net/http"
)

func (cfg *apiConfig) handlerAdminMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	htmlTemplate := fmt.Sprintf("<html>\n <body>\n <h1>Welcome, Chirpy Admin</h1>\n <p>Chirpy has been visited %d times!</p>\n </body>\n </html>", cfg.fileServerHits.Load())
	io.WriteString(w, htmlTemplate)
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)

		next.ServeHTTP(w, r)
	})
}
