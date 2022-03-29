package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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

type response struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

func getCharacters(c *gin.Context) {

	bodyBytes := exec("character", "GET")

	var responseObject response
	json.Unmarshal(bodyBytes, &responseObject)

	c.IndentedJSON(http.StatusOK, responseObject)

}

func exec(path string, method string) []byte {
	fmt.Println("Calling API...")

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

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes
}

func main() {

	router := gin.Default()
	router.GET("/character", getCharacters)

	router.Run("localhost:8080")

}
