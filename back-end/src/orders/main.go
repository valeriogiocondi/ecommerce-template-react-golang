package main

import (
	"fmt"
	"log"
	"net/http"

	Database "orders/database"
	Gateway "orders/gateway"
	env "orders/utils/env"
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
		log.Fatalf("The server has started with error %s\n", err)
	}
}
