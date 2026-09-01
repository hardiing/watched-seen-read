package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/hardiing/watched-seen-read/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	secret         string
	api_key        string
}

// This directive embeds the contents of the Vue dist folder into the variable
//
//go:embed frontend/dist/*
var frontendAssets embed.FS

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("JWT_SECRET")
	apiKey := os.Getenv("POLKA_KEY")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}
	dbQueries := database.New(db)
	const filepathRoot = "."
	const port = "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
		secret:         secret,
		api_key:        apiKey,
	}

	publicFS, err := fs.Sub(frontendAssets, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	frontendHandler := http.FileServer(http.FS(publicFS))

	fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	wrappedHandler := cfg.middlewareMetricsInc(fileServerHandler)

	mux := http.NewServeMux()
	mux.Handle("/app/", wrappedHandler)
	mux.Handle("/", frontendHandler)
	mux.HandleFunc("GET /api/healthz", readinessHandler)
	mux.HandleFunc("GET /admin/metrics", cfg.metricsHandler)
	mux.HandleFunc("POST /api/records", cfg.addRecordHandler)
	mux.HandleFunc("GET /api/records", cfg.getRecordsHandler)
	mux.HandleFunc("DELETE /api/records/{recordID}", cfg.deleteRecordHandler)

	corsHandler := enableCORS(mux)

	srv := &http.Server{
		Handler: corsHandler,
		Addr:    ":" + port,
	}
	log.Printf("Serving files from %s on port %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, WSR Admin</h1>
    <p>WSR has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
