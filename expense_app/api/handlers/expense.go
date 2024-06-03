// api/handlers/expenses.go
package handlers

import (
	"net/http"
	"encoding/json"
	"myapp/api/models"
	"time"
    "strconv"
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

    // create primary key for the new expense and created time
    newExpense := models.Expense{
        OrderNumber: orderNumber,
        ExpenseName: newExpenseRequest.ExpenseName,
        Expense:     newExpenseRequest.Expense,
        ExpenseDate: newExpenseRequest.ExpenseDate,
        RecordCreated: time.Now(),
    }
    expenses[orderNumber] = newExpense
    orderNumber++
    json.NewEncoder(w).Encode(newExpense)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
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
}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
    var allExpenses []models.Expense
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