package main

import (
	"net/http"
	"prototype/router"
)

func main() {
	mainHandler := router.ServerHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: &mainHandler,
	}
	server.ListenAndServe()
}
