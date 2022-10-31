package models_test

import (
	"context"
	"drp/logger/controllers"
	"drp/logger/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const testObjID string = "635f292a8f670ef141f609f4"

func TestEventLogValidation(t *testing.T) {
	tags := make([]string, 2)
	tags[0] = "app"
	tags[1] = "test"

	successRecord := models.EventLog{
		User:    "test-user",
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
		AppField: models.AppField{
			App: "test-app",
		},
	}
	successRecord.SetOnCreate()

	validate := validator.New()
	err := validate.Struct(successRecord)
	if err != nil {
		t.Errorf("Error initiating a new event log: %v", err)
	}

	noUserRecord := models.EventLog{
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
		AppField: models.AppField{
			App: "test-app",
		},
	}

	noUserRecord.SetOnCreate()

	validate = validator.New()
	err = validate.Struct(noUserRecord)
	if err == nil {
		t.Errorf("Unexpected validation pass for noUserRecord: %v", noUserRecord)
	}

	wrongDtRecord := models.EventLog{
		User:    "test-user",
		Dt:      -1,
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
		AppField: models.AppField{
			App: "test-app",
		},
	}

	validate = validator.New()
	err = validate.Struct(wrongDtRecord)
	if err == nil {
		t.Errorf("Unexpected validation pass for wrongDtRecord: %v", wrongDtRecord)
	}
}

func TestEventSingleCreate(t *testing.T) {
	tags := make([]string, 2)
	tags[0] = "app"
	tags[1] = "test"

	objectID, _ := primitive.ObjectIDFromHex(testObjID)

	eventLog := models.EventLog{
		Id:      objectID,
		User:    "test-user",
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
	}

	eventLogJson, _ := json.Marshal(eventLog)

	req, err := http.NewRequest("POST", "/event-log", strings.NewReader(string(eventLogJson)))
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.Create)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusCreated, status)
	}
}

func TestEventBulkCreate(t *testing.T) {
	tags := make([]string, 2)
	tags[0] = "app"
	tags[1] = "test"

	eventLogs := make([]models.EventLog, 2)

	eventLogs[0] = models.EventLog{
		User:    "test-user-1",
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
	}
	eventLogs[1] = models.EventLog{
		User:    "test-user-2",
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
	}

	eventLogsJson, _ := json.Marshal(eventLogs)

	req, err := http.NewRequest("POST", "/event-log/bulk", strings.NewReader(string(eventLogsJson)))
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.CreateBulk)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusCreated, status)
	}
}

func TestEventUpdate(t *testing.T) {
	tags := make([]string, 3)
	tags[0] = "app"
	tags[1] = "test"
	tags[2] = "update"

	eventLog := models.EventLog{
		User:    "test-user-update",
		Dt:      time.Now().Unix(),
		Event:   "User:login",
		Context: bson.M{},
		Tags:    tags,
	}

	eventLogJson, _ := json.Marshal(eventLog)

	req, err := http.NewRequest("PUT", "/event-log/"+testObjID, strings.NewReader(string(eventLogJson)))
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.Update)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusOK, status)
	}
}

func TestEventDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/event-log/"+testObjID, nil)
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.Delete)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusOK, status)
	}
}

func TestEventLogListBasic(t *testing.T) {
	req, err := http.NewRequest("GET", "/event-log", nil)
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.List)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusOK, status)
	}

	var response map[string]interface{}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
}

func TestEventLogListWithQuery(t *testing.T) {
	data := url.Values{}
	data.Set("page", "1")
	data.Set("pageSize", "10")
	data.Set("sort", "dt:DESC")

	req, err := http.NewRequest("GET", "/event-log", strings.NewReader(data.Encode()))
	ctx := context.WithValue(req.Context(), "app", "testing-app")
	req = req.WithContext(ctx)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.EventLogController{}.List)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong HTTP response code: expected: %d, got: %d.", http.StatusOK, status)
	}

	var response map[string]interface{}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
}
