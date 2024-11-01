package api

import (
	"database/sql"
	"log"
	"net/http"

	temp "github.com/Dariocent/time-checker-go/services/time-checker"
	"github.com/Dariocent/time-checker-go/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	// handle api/v1 with strip prefix
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", http.HandlerFunc(temp.Handler)))

	//Users
	// Initialize the store
	userStore := &user.SQLUserStore{DB: s.db}
	userService := user.UserService{UserStore: userStore}
	userHandler := user.UserHandler{UserService: userService}

	// Handle the routes
	// http.HandleFunc("GET /posts/{id}", handlePost2)
	router.HandleFunc("POST /api/v1/users", userHandler.CreateUserHandler)
	router.HandleFunc("GET /api/v1/users", userHandler.ListUsersHandler)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server is running on port: %s", s.addr)
	return server.ListenAndServe()
}
