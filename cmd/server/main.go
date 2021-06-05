package main

import (
	"os"

	"github.com/Sharykhin/go-assignment/server"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	server.ListenAndServe(port)
}
