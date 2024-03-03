package helper

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/lib/pq"
)

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)
		// Change Field Name
		switch match[1] {
		case "msisdn":
			match[1] = "phone number"
		}

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
		}
	}

	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

// CommonError is
func CommonError(err error) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
