package main

import (
	"log"
	"net/http"

	"almoxarifado-backend/config"
	"almoxarifado-backend/internal/database"
	"almoxarifado-backend/internal/handlers"
	"almoxarifado-backend/internal/repositories"
	"almoxarifado-backend/internal/services"
)

func main() {
	// Load config
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer db.Close()

	log.Println("🐘 connected to postgres")

	// Init layers
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	// Router
	mux := http.NewServeMux()

	// Health checks
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("API OK"))
	})

	mux.HandleFunc("/health/db", func(w http.ResponseWriter, _ *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("DB DOWN"))
			return
		}
		w.Write([]byte("DB OK"))
	})

	// Items endpoints
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			itemHandler.Create(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Start server
	log.Printf("🚀 API running on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, mux))
}
