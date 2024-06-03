package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bakageddy/gpa/templates"
)

const CSS_PATH = "index.css"

func css(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(CSS_PATH)
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	w.Header().Add("Content-type", "text/css")
	w.Write(data)
}

func home(w http.ResponseWriter, r *http.Request) {
	comp := templates.Home(CSS_PATH)
	comp.Render(context.Background(), w)
}

func form(w http.ResponseWriter, r *http.Request) {
	comp := templates.Semester(1, 3)
	comp.Render(context.Background(), w)
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	if r.ParseForm() != nil {
		log.Println("Failed to parse form!")
		return
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	comp := templates.Semester(1, 3)
	comp.Render(context.Background(), w)
}

func main() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/index.css", css)
	http.HandleFunc("/form", form)
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}
