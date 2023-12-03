package main

import (
	"log"
	"si-community/rest"
)

// @title Your Gin API
// @version 1.0
// @description This is a sample Gin API with Swagger documentation.
// @host localhost:8000
// @BasePath /v1
func main() {
	log.Println("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
