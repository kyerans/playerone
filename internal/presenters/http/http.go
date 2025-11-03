package httpx

import (
	"log"
	"net/http"

	"github.com/kyerans/playerone/internal/presenters/http/handlers"
)

func NewServer(hdl *handlers.Handler) *Server {
	return &Server{
		handler: hdl,
	}
}

type Server struct {
	handler *handlers.Handler
}

func (s *Server) ListenAndServe(addr string) error {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /license", s.handler.License)
	mux.HandleFunc("POST /license/release", s.handler.LicenseRelease)
	mux.HandleFunc("POST /license/register", s.handler.Register)

	log.Printf("[server] listen on: %s", addr)
	return http.ListenAndServe(addr, mux)
}
