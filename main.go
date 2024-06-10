package main

import (
	"log"
	"net/http"

	"github.com/bakageddy/gpa/services"
)

func main() {
	http.HandleFunc("/cgpa", services.Cgpa)
	http.HandleFunc("/subject/add", services.SubjectAdd)
	http.HandleFunc("/subject/delete", services.SubjectDelete)
	http.HandleFunc("/semester/add", services.SemesterAdd)
	http.HandleFunc("/semester/delete", services.SemesterDelete)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", services.Home)

	log.Println("Starting Server!")
	http.ListenAndServe(":80", nil)
	log.Println("Ending Server!")
}
