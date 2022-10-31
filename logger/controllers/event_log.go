package controllers

import (
	"drp/logger/models"
	"drp/logger/services/entity"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type EventLogController struct{}

func (ec EventLogController) List(w http.ResponseWriter, r *http.Request) {
	result, err := entity.List[*models.EventLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}

func (ec EventLogController) Create(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Create[*models.EventLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusCreated,
		"data":       result,
	})

}

func (ec EventLogController) CreateBulk(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.CreateBulk[*models.EventLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusCreated,
		"data":       result,
		"errors":     validationErrors,
	})
}

func (ec EventLogController) Update(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Update[*models.EventLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
		"errors":     validationErrors,
	})
}

func (ec EventLogController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Get[*models.EventLog](r)

	httpCode := http.StatusUnprocessableEntity
	if err == mongo.ErrNoDocuments {
		httpCode = http.StatusNotFound
	}

	processError(w, err, httpCode)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}

func (ec EventLogController) Delete(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Delete[*models.EventLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}
