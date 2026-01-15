package main

import (
	"almoxarifado-backend/config"
	"almoxarifado-backend/internal/database"
	"almoxarifado-backend/internal/handlers"
	"almoxarifado-backend/internal/middleware"
	"almoxarifado-backend/internal/repositories"
	"almoxarifado-backend/internal/services"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Carregando configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("erro ao carregar configuração:", err) // Coloquei isso pq se faltar JWT ouDB password o app não inicia
	}

	// Conexão com o banco de dados
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer db.Close()

	log.Println("connected to postgres")

	// Executando migrations
	if err := database.RunMigrations(db, cfg.DatabaseURL()); err != nil {
		log.Fatal("erro ao executar migrations:", err)
	}
	log.Println("migrations executed")

	// Seed admin padrão
	if err := database.SeedAdminUser(db, cfg.AdminDefaultEmail, cfg.AdminDefaultPassword); err != nil {
		log.Fatal("erro ao criar admin padrão:", err)
	}

	// Inicializando repositórios, serviços e handlers
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	// auth
	adminRepo := repositories.NewAdminRepository(db)
	authService := services.NewAuthService(adminRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	// Rotas
	mux := http.NewServeMux()

	// Verificação de saúde
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

	// Auth com Rate Limiter
	loginMux := http.NewServeMux()
	loginMux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authHandler.Login(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	mux.Handle("/auth/login", middleware.RateLimitMiddleware(loginMux))

	// Items (protegido)
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			itemHandler.Create(w, r)
			return
		}
		if r.Method == http.MethodGet {
			itemHandler.GetAll(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Rota para obter item por ID
	protectedMux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			itemHandler.GetByID(w, r)
			return
		}
		if r.Method == http.MethodPut {
			itemHandler.Update(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			itemHandler.Delete(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Aplicar middleware de autenticação
	mux.Handle("/items", middleware.AuthMiddleware(authService)(protectedMux))
	mux.Handle("/items/", middleware.AuthMiddleware(authService)(protectedMux))

	//CORS Middleware
	corsHandler := cors.New(cors.Options{
		// Lembrar de quando for fazer o front com Vue, colocar o link certo no Origins
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	// Iniciando o servidor
	log.Printf("API running on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, corsHandler))
}
