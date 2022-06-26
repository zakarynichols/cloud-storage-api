package http

import (
	video_stuff "cloud-storage-api"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router

	Addr string

	UploadService video_stuff.UploadService
}

func NewServer() *Server {
	router := mux.NewRouter()

	s := &Server{
		server: &http.Server{
			Handler:      router,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		router: router,
	}

	// Custom not found handler
	// s.router.NotFoundHandler = http.HandlerFunc(s.notFoundHandler)
	// Custom method not allowed handler
	// s.router.MethodNotAllowedHandler = http.HandlerFunc(s.methodNotAllowedHandler)

	// middleware
	s.router.Use(s.loggingHandler)

	// routes
	s.registerUploadRoutes(s.router)

	return s
}

func (s *Server) ListenAndServe() error {
	var err error

	s.ln, err = net.Listen("tcp", s.Addr)

	if err != nil {
		return err
	}

	return s.server.Serve(s.ln)
}

// TODO: expand on loggingHandler
func (s *Server) loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
