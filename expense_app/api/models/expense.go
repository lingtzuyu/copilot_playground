// api/models/expense.go
package models

import "time"

type Expense struct {
	OrderNumber     int
	ExpenseName     string
	Expense 	    int
	ExpenseDate     time.Time
	RecordCreated   time.Time
}

type ExpenseRequest struct {
	ExpenseName     string
	Expense 	    int
	ExpenseDate     time.Time
}