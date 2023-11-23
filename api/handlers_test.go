package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"zocket-api/db"
	"zocket-api/models"
)

func TestCreateUser(t *testing.T) {

	db.Connect()

	// Create a new user
	user := models.User{
		ID:        [12]byte{},
		Name:      "John",
		Mobile:    "1234567890",
		Latitude:  "12.345",
		Longitude: "67.890",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	// Convert user struct to JSON
	userJSON, _ := json.Marshal(user)

	// Create a request with the user JSON
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create an HTTP handler function for CreateUser
	handler := http.HandlerFunc(CreateUser)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	expected := `{"id":"000000000000000000000000","name":"John Doe","mobile":"1234567890","latitude":"12.345","longitude":"67.890","created_at":"` + time.Now().UTC().Format(time.RFC3339) + `","updated_at":"` + time.Now().UTC().Format(time.RFC3339) + `"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
