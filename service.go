package main

import (
	"fmt"
	"time"
)

func taskTimestamp() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func nextExpenseID(expenses []Expense) int {
	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}
	return maxID + 1
}

func addExpense(expenses []Expense, amount float64, category string) []Expense {
	timestamp := taskTimestamp()
	newExpense := Expense{
		Amount:   amount,
		Category: category,
		Date:     timestamp.Format(time.RFC3339),
		Note:     "",
		ID:       nextExpenseID(expenses),
	}
	return append(expenses, newExpense)
}

func listExpenses(expenses []Expense) {
	printExpenseHeader()
	for _, expense := range expenses {
		printExpenseRow(expense)
	}
}

func printExpenseHeader() {
	fmt.Printf("%-4s %-12s %-16s %-16s %-16s\n", "ID", "AMOUNT", "NOTE", "CATEGORY", "DATE")
}

func printExpenseRow(expense Expense) {
	fmt.Printf("%-4d %-12.2f %-16s %-16s %-16s\n", expense.ID, expense.Amount, expense.Note, expense.Category, expense.Date)
}
