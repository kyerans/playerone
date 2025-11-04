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

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Range, Accept, Content-Type, Origin")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")

		// Cho phép preflight request (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // hoặc cụ thể: http://localhost:3000
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h(w, r)
	}
}

func (s *Server) ListenAndServe(addr string) error {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /license", cors(s.handler.License))
	mux.HandleFunc("GET /license", cors(s.handler.GetLicense))
	mux.HandleFunc("POST /license/release", cors(s.handler.LicenseRelease))
	mux.HandleFunc("POST /license/register", cors(s.handler.Register))

	fs := http.FileServer(http.Dir("./media/hls"))
	mux.Handle("/hls/", withCORS(http.StripPrefix("/hls/", fs)))

	log.Printf("[server] listen on: %s", addr)
	return http.ListenAndServe(addr, mux)
}
