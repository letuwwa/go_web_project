package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name      string
	age       uint16
	money     int16
	avgGrades float64
	level     float64
}

func homePage(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Go is quite easy")
}

func contacts(w http.ResponseWriter, request *http.Request) {
	testUser := User{
		name:      "John",
		age:       22,
		money:     200,
		avgGrades: 12.3,
		level:     22.3,
	}
	fmt.Fprintf(w, "User is "+testUser.name)
}

func handleServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contacts)
	httpError := http.ListenAndServe("localhost:8081", nil)
	if httpError != nil {
		fmt.Println(httpError)
	}
}

func main() {
	handleServer()
}
