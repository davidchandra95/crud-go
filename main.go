package main

import (
	"crud-go/config"
	"crud-go/handler"
	"crud-go/middleware"
	"crud-go/repository"
	"crud-go/service"
	"crud-go/usecase"
	"crud-go/util"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database connection
	db, err := sql.Open("postgres", cfg.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Run database migrations
	err = util.RunMigrations(db, "migrations")
	if err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	// Uncomment this line to rollback the last migration
	// err = util.RollbackLastMigration(db, "migrations")
	// if err != nil {
	// 	log.Fatalf("Could not rollback migration: %v", err)
	// }

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)

	// Initialize usecases
	userUsecase := usecase.NewUserUsecase(userService)
	authUsecase := usecase.NewAuthUsecase(authService)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Setup router and routes
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/signup", authHandler.SignUp).Methods("POST")
	r.HandleFunc("/signin", authHandler.SignIn).Methods("POST")

	// Protected routes
	authMiddleware := middleware.AuthMiddleware(cfg.JWTSecret)
	r.Handle("/users", authMiddleware(http.HandlerFunc(userHandler.GetUsers))).Methods("GET")
	r.Handle("/users/{id:[0-9]+}", authMiddleware(http.HandlerFunc(userHandler.GetUserByID))).Methods("GET")
	r.Handle("/users", authMiddleware(http.HandlerFunc(userHandler.CreateUser))).Methods("POST")
	r.Handle("/users/{id:[0-9]+}", authMiddleware(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")
	r.Handle("/users/{id:[0-9]+}", authMiddleware(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
