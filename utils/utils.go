package utils

import (
	"net/url"
	"strconv"
	"errors"
	"github.com/bakageddy/gpa/types"
)

func CalculateGPA(f url.Values) (types.GPA, error) {
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

		if (creditval < 0) {
			return types.GPA{}, errors.New("Malformed Request")
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
