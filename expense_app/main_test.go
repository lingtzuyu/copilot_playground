package main

import (
	"bytes"
	"encoding/json"
	"myapp/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	// Create a request body
	expense := map[string]interface{}{
		"ExpenseName": "Test Expense",
		"Expense":     100.0,
		"ExpenseDate": "2024-01-01T00:00:00Z",
		"Category":    "food",
	}
	body, _ := json.Marshal(expense)

	// Create a request
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router like in your main function
	r := mux.NewRouter()
	r.HandleFunc("/expenses", handlers.CreateExpense).Methods("POST")

	// Serve the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Test Expense", response["ExpenseName"])
	assert.Equal(t, 100.0, response["Expense"])
	assert.Equal(t, "2024-01-01T00:00:00Z", response["ExpenseDate"])
	assert.Equal(t, "food", response["Category"])
}

func TestDeleteExpense(t *testing.T) {
	// First, create a new expense
	expense := map[string]interface{}{
		"ExpenseName": "Test Expense",
		"Expense":     100.0,
		"ExpenseDate": "2024-01-01T00:00:00Z",
		"Category":    "food",
	}
	body, _ := json.Marshal(expense)
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/expenses", handlers.CreateExpense).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response to get the ID of the created expense
	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	id := response["id"].(string)

	// Now, delete the expense
	req, _ = http.NewRequest("DELETE", "/expenses/"+id, nil)
	rr = httptest.NewRecorder()
	r.HandleFunc("/expenses/{id}", handlers.DeleteExpense).Methods("DELETE")
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Add more assertions based on the expected response
}
