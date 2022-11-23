package exception

import (
	"go-rest/entity/web"
	"go-rest/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundErr(w, r, err) {
		return
	}

	if validationErr(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErr(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webRes := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResBody(w, webRes)
		return true
	} else {
		return false
	}
}

func notFoundErr(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundErr)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webRes := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResBody(w, webRes)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webRes := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResBody(w, webRes)
}
