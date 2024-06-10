// api/handlers/expenses.go
package handlers

import (
	"encoding/json"
	"myapp/api/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// make a in-memory store for expenses, start from 1
var expenses = make(map[int]models.Expense)
var orderNumber = 1

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var newExpenseRequest models.ExpenseRequest
	err := json.NewDecoder(r.Body).Decode(&newExpenseRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the Expense is a valid number
	if newExpenseRequest.Expense <= 0 {
		http.Error(w, "Invalid expense amount", http.StatusBadRequest)
		return
	}

	// Check if the Category is valid
	validCategories := []string{"food", "clothing", "housing", "transportation"}
	if !contains(validCategories, newExpenseRequest.Category) {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	// Check if the ExpenseDate is not more than a year ago
	if newExpenseRequest.ExpenseDate.Before(time.Now().AddDate(-1, 0, 0)) {
		http.Error(w, "ExpenseDate cannot be more than a year ago", http.StatusBadRequest)
		return
	}

	// create primary key for the new expense and created time
	newExpense := models.Expense{
		OrderNumber:   orderNumber,
		ExpenseName:   newExpenseRequest.ExpenseName,
		Expense:       newExpenseRequest.Expense,
		ExpenseDate:   newExpenseRequest.ExpenseDate,
		RecordCreated: time.Now(),
		Category:      newExpenseRequest.Category,
	}
	expenses[orderNumber] = newExpense
	orderNumber++
	json.NewEncoder(w).Encode(newExpense)
}

/* func GetExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderNumber, err := strconv.Atoi(params["orderNumber"])
	if err != nil {
		http.Error(w, "Invalid order number", http.StatusBadRequest)
		return
	}

	if expense, exists := expenses[orderNumber]; exists {
		json.NewEncoder(w).Encode(expense)
	} else {
		http.Error(w, "Expense not found", http.StatusNotFound)
	}
} */

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	allExpenses := make([]models.Expense, 0) // Initialize as empty slice
	for _, expense := range expenses {
		allExpenses = append(allExpenses, expense)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allExpenses)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderNumber, err := strconv.Atoi(params["orderNumber"])
	if err != nil {
		http.Error(w, "Invalid order number", http.StatusBadRequest)
		return
	}

	if _, exists := expenses[orderNumber]; exists {
		delete(expenses, orderNumber)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Expense not found", http.StatusNotFound)
	}
}

// Helper function to check if a slice contains a string
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func SearchExpenses(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	expenseName := r.URL.Query().Get("expenseName")
	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	var startDate, endDate time.Time
	var err error

	// Parse dates if they are provided
	if startDateStr != "" {
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			http.Error(w, "Invalid start date", http.StatusBadRequest)
			return
		}
	}
	if endDateStr != "" {
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			http.Error(w, "Invalid end date", http.StatusBadRequest)
			return
		}
	}

	// Filter expenses
	filteredExpenses := make([]models.Expense, 0)
	for _, expense := range expenses {
		if expenseName != "" && !strings.Contains(expense.ExpenseName, expenseName) {
			continue
		}
		if !startDate.IsZero() && expense.ExpenseDate.Before(startDate) {
			continue
		}
		if !endDate.IsZero() && expense.ExpenseDate.After(endDate) {
			continue
		}
		filteredExpenses = append(filteredExpenses, expense)
	}

	// Return filtered expenses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredExpenses)
}
