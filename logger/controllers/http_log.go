package controllers

import (
	"drp/logger/models"
	"drp/logger/services/entity"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type HttpLogController struct{}

func (ec HttpLogController) List(w http.ResponseWriter, r *http.Request) {
	result, err := entity.List[*models.HttpLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}

func (ec HttpLogController) Create(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Create[*models.HttpLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusCreated,
		"data":       result,
	})

}

func (ec HttpLogController) CreateBulk(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.CreateBulk[*models.HttpLog](r)

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

func (ec HttpLogController) Update(w http.ResponseWriter, r *http.Request) {
	result, validationErrors, err := entity.Update[*models.HttpLog](r)

	processValidationErrors(w, validationErrors)
	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
		"errors":     validationErrors,
	})
}

func (ec HttpLogController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Get[*models.HttpLog](r)

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

func (ec HttpLogController) Delete(w http.ResponseWriter, r *http.Request) {
	result, err := entity.Delete[*models.HttpLog](r)

	processError(w, err, http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": http.StatusOK,
		"data":       result,
	})
}
