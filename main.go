package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPAge)
	server := ":8080"
	fmt.Println("Server Listen on port", server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}

}

func usersPAge(w http.ResponseWriter, r *http.Request) {
	users := []User{{"Ali", 77}, User{"Human", 88}}

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}
}
