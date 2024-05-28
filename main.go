package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/bakageddy/gpa/templates"
)

func id_generator() func() int {
	id := 0
	return func() int {
		id += 1
		return id
	}
}

var id func() int = id_generator()

func add_subject(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Subject()
	comp.Render(context.Background(), w)
}

func calculate_gpa(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if (err != nil) {
		w.Write([]byte("Error calculating GPA"))
	}

	log.Println(string(body))
	w.Write([]byte("hehe boi"))
}

func add_semester(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Semester(id())
	comp.Render(context.Background(), w)
}

func home(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Home()
	comp.Render(context.Background(), w)
}

func main() {
	http.HandleFunc("/calculate_gpa", calculate_gpa)
	http.HandleFunc("/add-sem", add_semester)
	http.HandleFunc("/add-sub", add_subject)
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}
