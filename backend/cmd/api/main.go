package main

import (
	"log"
	"net/http"

	"almoxarifado-backend/config"
	"almoxarifado-backend/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("API OK"))
	})

	dbHealthHandler := func(w http.ResponseWriter, _ *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("DB DOWN"))
			return
		}
		w.Write([]byte("DB OK"))
	}

	mux.HandleFunc("/health/db", dbHealthHandler)
	mux.HandleFunc("/health/db/", dbHealthHandler)

	log.Printf("API running on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, mux))
}
