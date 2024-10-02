package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			_ = WriteJson(w, http.StatusBadRequest, Error{Error: err.Error()})
		}
	}
}

func StartServer(listenPort string) ServerInterface {
	return &Server{
		listenPort: listenPort,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/access", makeHTTPHandlerFunc(s.handleAccess))
	router.HandleFunc("/access/{prof_id}", makeHTTPHandlerFunc(s.handleAccess))

	log.Printf("Started server at port: %s", s.listenPort)
	if err := http.ListenAndServe(s.listenPort, router); err != nil {
		log.Fatal(err)
	}
}
