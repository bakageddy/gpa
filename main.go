package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bakageddy/gpa/services"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cgpa", services.Cgpa)
	mux.HandleFunc("/gpa", services.Gpa)
	mux.HandleFunc("/subject/add", services.SubjectAdd)
	mux.HandleFunc("/subject/delete", services.SubjectDelete)
	mux.HandleFunc("/semester/add", services.SemesterAdd)
	mux.HandleFunc("/semester/delete", services.SemesterDelete)
	mux.HandleFunc(fmt.Sprintf("/%s", services.CSS_PATH), services.CSS)
	mux.HandleFunc("/", services.Home)

	log.Println("Starting Server!")
	http.ListenAndServe(":8080", mux)
	log.Println("Ending Server!")
}
