package server

import (
	"fmt"
	"net/http"
	"rick-morty/api/data"
)

func (s *Server) ListAll(rw http.ResponseWriter, hr *http.Request) {
	s.l.Debug("Get all records")

	rw.Header().Add("Accept", "application/json")
	rw.Header().Add("Content-Type", "application/json")

	res, err := s.response.Execute("character", http.MethodGet, rw)
	if err != nil {
		fmt.Print(err.Error())
	}

	data.ToJSON(res, rw)
}
