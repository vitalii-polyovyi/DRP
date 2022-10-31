package controllers

import (
	"drp/logger/models"
	"drp/logger/services/entity"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type EntityLogController struct{}

func (ec EntityLogController) List(w http.ResponseWriter, r *http.Request) {
	result, err := entity.List[*models.EntityLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}

func (ec EntityLogController) Create(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Create[*models.EntityLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusCreated,
		"data":       result,
	})

}

func (ec EntityLogController) CreateBulk(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.CreateBulk[*models.EntityLog](r)

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

func (ec EntityLogController) Update(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Update[*models.EntityLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
		"errors":     validationErrors,
	})
}

func (ec EntityLogController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Get[*models.EntityLog](r)

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

func (ec EntityLogController) Delete(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Delete[*models.EntityLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}
