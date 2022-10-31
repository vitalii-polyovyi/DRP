package controllers

import (
	"drp/logger/helpers"
	"encoding/json"
	"log"
	"net/http"
)

func processError(w http.ResponseWriter, err error, httpCode int) {
	if err != nil {
		w.WriteHeader(httpCode)
		b, _ := json.Marshal(&helpers.ErrorResponse{Error: err.Error()})
		w.Write(b)

		log.Panicf("%v", err)
	}
}

func processValidationErrors(w http.ResponseWriter, err []*helpers.ErrorResponse) {
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		b, _ := json.Marshal(&err)
		w.Write(b)

		log.Panicf("%v", err)
	}
}
