// api/models/expense.go
package models

import "time"

type Expense struct {
	OrderNumber   int
	ExpenseName   string
	Expense       int
	ExpenseDate   time.Time
	RecordCreated time.Time
	Category      string
}

type ExpenseRequest struct {
	ExpenseName string
	Expense     int
	ExpenseDate time.Time
	Category    string
}
