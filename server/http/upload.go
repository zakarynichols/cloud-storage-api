package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 10 MB max limit for now
const MAX_UPLOAD_SIZE = 1024 * 1024 * 10

func (s *Server) registerUploadRoutes(r *mux.Router) {
	r.HandleFunc("/upload", s.upload)
}

// Upload endpoint
func (s *Server) upload(w http.ResponseWriter, r *http.Request) {
	// Read header values
	origin := r.Header.Get("Origin")
	log.Print(origin)

	if origin == "" {
		log.Fatal("missing origin header")
	}

	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	// Prevent a file upload larger than preset
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	// The key provided to `FormFile` must match the `name` attribute on the HTML form/input.
	file, handler, err := r.FormFile("upload")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err = s.UploadService.Upload(ctx, file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ctx.Done()
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(200)

	json.NewEncoder(w).Encode(struct {
		Filename string `json:"filename"`
		Size     int64  `json:"size"`
	}{Filename: handler.Filename, Size: handler.Size})
}
