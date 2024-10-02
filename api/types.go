package api

import "net/http"

type ServerInterface interface {
	Run()
	handleAccess(http.ResponseWriter, *http.Request) error
}

type Server struct {
	listenPort string
}

type Error struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

// api "/access"
type AccessRequest struct {
	ProfileIds interface{} `json:"profile_id"`
}
type AccessResponseElement struct {
	ProfileId  int64  `json:"profile_id"`
	Accessible string `json:"is_accessible"`
}
type AccessResponse []AccessResponseElement
