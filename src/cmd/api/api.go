package api

import (
	"database/sql"
	"log"
	"net/http"

	temp "github.com/Dariocent/time-checker-go/services/time-checker"
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

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server is running on port: %s", s.addr)
	return server.ListenAndServe()
}
