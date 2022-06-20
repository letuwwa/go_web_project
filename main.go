package main

import (
	"go_web_project/pkg/db"
	"log"
	"net/http"
)

func main() {
	db.Init()

	log.Println("API is running!")
	http.ListenAndServe(":4000", nil)
}
