package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenPort string
}

func StartServer(listenPort string) *Server {
	return &Server{
		listenPort: listenPort,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	log.Printf("Started server at port: %s", s.listenPort)
	if err := http.ListenAndServe(s.listenPort, router); err != nil {
		log.Fatal(err)
	}
}
