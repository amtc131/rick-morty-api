package data

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	URL     string   `json:"url"`
	Created string   `json:"created"`
}

type Response struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

type ResponseBD struct {
	log hclog.Logger
}

func NewResponseDB(l hclog.Logger) *ResponseBD {
	return &ResponseBD{l}
}

func (r *ResponseBD) Execute(path string, method string) (*http.Response, error) {
	r.log.Info("Calling API...")

	endPoint := fmt.Sprintf("https://rickandmortyapi.com/api/%s", path)

	client := &http.Client{}

	req, err := http.NewRequest(method, endPoint, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	return resp, err
}
