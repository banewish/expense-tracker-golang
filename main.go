package main

import (
	"time"
	//"github.com/lib/pq"
)

func main() {

	listExpenses([]Expense{
		{ID: 1, Amount: 50.0, Category: "Food", Note: "Lunch", Date: time.Now().Format(time.RFC3339)},
		{ID: 2, Amount: 20.0, Category: "Transport", Note: "Bus fare", Date: time.Now().Format(time.RFC3339)},
		{ID: 3, Amount: 100.0, Category: "Shopping", Note: "Clothes", Date: time.Now().Format(time.RFC3339)},
	})
}
