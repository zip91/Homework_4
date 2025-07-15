package main

import (
	"log"
	"todo/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Server started at http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
