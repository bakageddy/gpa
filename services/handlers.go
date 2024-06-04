package services

import (
	"log"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/bakageddy/gpa/templates"
	"github.com/bakageddy/gpa/utils"
)

const CSS_PATH string = "assets/css/index.css"

func Home(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Home(CSS_PATH)
	comp.Render(context.Background(), w)
}

func Cgpa(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		log.Println("Failed to parse form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	templates.CGPA(0).Render(context.Background(), w)
}

func Gpa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	if r.ParseForm() != nil {
		log.Println("Failed to parse form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	score, err := utils.CalculateGPA(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	templates.GPA(score).Render(context.Background(), w)
}

func SubjectAdd(w http.ResponseWriter, _ *http.Request) {
	templates.Subject().Render(context.Background(), w)
}

func SubjectDelete(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func SemesterAdd(w http.ResponseWriter, _ *http.Request) {
	templates.Semester(0).Render(context.Background(), w)
}

func SemesterDelete(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CSS(w http.ResponseWriter, _ *http.Request) {
	file, err := os.Open(CSS_PATH)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "text/css")
	w.Write(body)
}
