package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"rick-morty/api/data"
	server "rick-morty/api/server"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIN_ADDRES", false, ":9093", "Bin addres for the server")

func main() {

	env.Parse()
	l := hclog.Default()

	db := data.NewResponseDB(l)
	ch := server.NewCharacter(l, db)

	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/character", ch.ListAll)

	s := http.Server{
		Addr:         *bindAddress,                                     // configure the bind addres
		Handler:      sm,                                               // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		IdleTimeout:  120 * time.Second,                                // max time to read request from the client
		ReadTimeout:  1 * time.Second,                                  // max time for connections using TCP keep-alive
		WriteTimeout: 1 * time.Second,                                  // max tie to write response to the client
	}

	//start the server
	go func() {
		l.Info("Starting server on port 9093")
		err := s.ListenAndServe()
		if err != nil {
			l.Info("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Info("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(ctx)
}
