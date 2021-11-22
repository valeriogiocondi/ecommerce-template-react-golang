package main

import (
	"fmt"
	"log"
	"net/http"

	Database "customers/database"
	Gateway "customers/gateway"
	env "customers/utils/env"
)

func main() {

	startServer()
}

func startServer() {

	port := env.ReadFile("PORT")

	fmt.Println("Server is starting...")

	Gateway.Init()
	Database.Init()
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("The server has started with error %s\n", err)
	}
}
