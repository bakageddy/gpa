package services

import (
	"log"
	"context"
	"net/http"

	"github.com/bakageddy/gpa/templates"
	"github.com/bakageddy/gpa/utils"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Home()
	comp.Render(context.Background(), w)
}

func Cgpa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	if r.ParseForm() != nil {
		log.Println("Failed to parse form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	templates.Semester().Render(context.Background(), w)
}

func SemesterDelete(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
