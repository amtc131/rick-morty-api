package main

import (
	"io/ioutil"
	"net/http"

	"github.com/amtc131/client/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	l := hclog.Default()

	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}

	l.Info("Calling api")

	httpClientApi := server.NewHttpClientApi("https://rickandmortyapi.com/api/character", &http.Client{})

	response, err := httpClientApi.
		Method(http.MethodGet).
		Headers(headers).
		Do()

	if err != nil {
		l.Error("Error to call api", err)
		return
	}

	//Close body response
	defer response.Body.Close()

	datBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.Error("Error to read data", err)
		return
	}
	l.Info(string(datBytes))
}
