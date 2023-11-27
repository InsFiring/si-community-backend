package main

import (
	"log"
	"si-community/rest"
)

func main() {
	log.Println("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
