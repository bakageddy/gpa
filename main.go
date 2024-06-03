package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/bakageddy/gpa/templates"
	"github.com/bakageddy/gpa/types"
)

const CSS_PATH string = "assets/css/index.css"

func home(w http.ResponseWriter, _ *http.Request) {
	comp := templates.Home(CSS_PATH)
	comp.Render(context.Background(), w)
}

func calculate_gpa(f url.Values) (types.GPA, error) {
	cred_list_s := f["subject_credit"]
	grade_list_s := f["subject_grade"]

	if len(cred_list_s) != len(grade_list_s) || len(cred_list_s) == 0 {
		return types.GPA{}, errors.New("Malformed Request")
	}

	total_subject_cred := 0
	total_cred := 0
	for i := range len(cred_list_s) {
		creditval, err := strconv.Atoi(cred_list_s[i])
		if err != nil {
			return types.GPA{}, err
		}
		total_subject_cred += creditval

		gradeval, err := strconv.Atoi(grade_list_s[i])
		if err != nil {
			return types.GPA{}, err
		}
		total_cred += gradeval * creditval
	}

	return types.GPA{
		TotalCredits:        total_cred,
		TotalSubjectCredits: total_subject_cred,
		Score:               float64(total_cred) / float64(total_subject_cred),
	}, nil
}

func gpa(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	score, err := calculate_gpa(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	templates.GPA(score).Render(context.Background(), w)
}

func subject_add(w http.ResponseWriter, _ *http.Request) {
	templates.Subject().Render(context.Background(), w)
}

func subject_delete(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func css(w http.ResponseWriter, _ *http.Request) {
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

func main() {
	http.HandleFunc("/subject/add", subject_add)
	http.HandleFunc("/subject/delete", subject_delete)
	http.HandleFunc(fmt.Sprintf("/%s", CSS_PATH), css)
	http.HandleFunc("/gpa", gpa)
	http.HandleFunc("/", home)

	log.Println("Starting Server!")
	http.ListenAndServe(":8080", nil)
	log.Println("Ending Server!")
}
