package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Server is a struct that represents the ServerConfig
type Server struct {
	Logger *log.Logger
}

type response struct {
	Status  int    `json:",omitempty"`
	Method  string `json:",omitempty"`
	Error   error  `json:",omitempty"`
	Message string `json:",omitempty"`
	Hash    string `json:",omitempty"`
	Path    string `json:",omitempty"`
}

// New creates a new server.
func New(log *log.Logger) Server {
	return Server{
		Logger: log,
	}
}

// LoggingMiddleware injects a debug logger
func (s Server) LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Printf(fmt.Sprintf("[DEBUG] %v %v %v", r.Method, r.RequestURI, r.Header.Get("Content-Type")))
		next.ServeHTTP(w, r)
	})

}

func (s Server) writeResponse(w http.ResponseWriter, res response) {

	w.WriteHeader(res.Status)
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(res)
	if res.Error != nil || err != nil {
		s.Logger.Printf(fmt.Sprintf("[ERROR] %v %v", res.Message, res.Error))
	}

	w.Write(b)
}
