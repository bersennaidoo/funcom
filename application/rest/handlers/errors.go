package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func ErrorMessages(err int64) (int, int, string) {
	errorMessage := ""
	statusCode := 200
	errorCode := 0
	switch err {
	case 1062:
		errorMessage = "Duplicate entry"
		errorCode = 10
		statusCode = http.StatusConflict
	}

	return errorCode, statusCode, errorMessage

}

func dbErrorParse(err string) (string, int64) {
	Parts := strings.Split(err, ":")
	errorMessage := Parts[1]
	Code := strings.Split(Parts[0], "Error ")
	errorCode, _ := strconv.ParseInt(Code[1], 10, 32)
	return errorMessage, errorCode
}
