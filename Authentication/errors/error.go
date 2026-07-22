package errors

import (
	"net/http"

	"github.com/alilacream/auth/models"
)

// the display function
func DisplayErr(w http.ResponseWriter, err string) {
	httpErr := Errors(w, err)
	if httpErr != nil {
		http.Error(httpErr.Writer, httpErr.ErrStat, int(httpErr.Status))
	}
}

// switch case for possible errors
func Errors(w http.ResponseWriter, err string) *models.HTTPError {
	switch err {
	case "UnAuthorized":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: err,
			Status:  http.StatusUnauthorized,
		})

	case "Method":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: "UnAuthorized Method",
			Status:  http.StatusMethodNotAllowed,
		})
	case "Parse":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: "Unknown Request Format",
			Status:  http.StatusBadRequest,
		})
	case "Password":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: "Incorect Password",
			Status:  http.StatusConflict,
		})
	case "Userregister":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: "User Already Exists",
			Status:  http.StatusConflict,
		})
	case "Userlogin":
		return new(models.HTTPError{
			Writer:  w,
			ErrStat: "Used Does not exist",
			Status:  http.StatusConflict,
		})
	}
	return nil
}
