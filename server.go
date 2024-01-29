package db

import "net/http"

type Server struct {
	Schema  *Schema
	RootDir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
