package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	Name  string
	Level int
}

var tmpl = template.Must(template.ParseFiles("badge.html"))

func badgeHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Name:  "John Doe",
		Level: 3,
	}

	if err := tmpl.Execute(w, student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", badgeHandler)
	http.ListenAndServe(":8080", nil)
}
