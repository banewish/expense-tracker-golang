package main

import (
	"database/sql"
	"fmt"
	"time"
)

func addExpense(title string, amount float64, category string) error {
	if title == "" || amount <= 0 {
		return fmt.Errorf("invalid input: title and positive amount required")
	}

	sqlStatement := `INSERT INTO expenses (title, amount, category, expense_date) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, title, amount, category, time.Now().Format("2006-01-02"))
	if err != nil {
		return fmt.Errorf("error inserting expense: %w", err)
	}
	return nil
}

func listExpenses() error {
	rows, err := db.Query(`SELECT id, title, amount, category, expense_date FROM expenses ORDER BY expense_date DESC`)
	if err != nil {
		return fmt.Errorf("error querying expenses: %w", err)
	}
	defer rows.Close()

	printExpenseHeader()
	hasRows := false
	for rows.Next() {
		hasRows = true
		var expense Expense
		var expenseDate sql.NullString
		err := rows.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Category, &expenseDate)
		if err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}
		if expenseDate.Valid {
			expense.Date = expenseDate.String
		}
		printExpenseRow(expense)
	}

	if !hasRows {
		fmt.Println("No expenses found. Add one with: expense-tracker add <title> <amount> <category>")
	}
	return rows.Err()
}

func printExpenseHeader() {
	fmt.Printf("%-4s %-20s %-12s %-16s %-12s\n", "ID", "TITLE", "AMOUNT", "CATEGORY", "DATE")
}

func printExpenseRow(expense Expense) {
	fmt.Printf("%-4d %-20s %-12.2f %-16s %-12s\n", expense.ID, expense.Title, expense.Amount, expense.Category, expense.Date)
}
