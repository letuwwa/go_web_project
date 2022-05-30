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

func (user User) getAllInfo() string {
	return fmt.Sprintf(
		"User name is %s. He is %d y.o. and he has $%d",
		user.name,
		user.age,
		user.money,
	)
}
func (user *User) setNewName(newName string) {
	user.name = newName
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Go is quite easy")
}
func contacts(writer http.ResponseWriter, request *http.Request) {
	testUser := User{
		name:      "John",
		age:       22,
		money:     200,
		avgGrades: 12.3,
		level:     22.3,
	}
	testUser.setNewName("test_user")
	fmt.Fprintf(writer, testUser.getAllInfo())
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
