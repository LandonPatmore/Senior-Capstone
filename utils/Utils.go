package utils

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Err struct {
	Error string
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, message string) {
	render.JSON(w, r, Err{Error: message})
}
