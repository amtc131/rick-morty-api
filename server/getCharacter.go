package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rick-morty/api/data"
)

func (s *Server) ListAll(rw http.ResponseWriter, hr *http.Request) {

	s.l.Info("Get all records ")

	rw.Header().Add("Accept", "application/json")
	rw.Header().Add("Content-Type", "application/json")

	resp, err := s.response.Execute("character", http.MethodGet)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	response := data.Response{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Print(err.Error())
	}

	data.ToJSON(response, rw)
}
