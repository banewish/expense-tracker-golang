package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	expenseDatabase()

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: expense-tracker add <title> <amount> <category>")
			os.Exit(1)
		}
		title := os.Args[2]
		amount := parseAmount(os.Args[3])
		category := os.Args[4]
		err := addExpense(title, amount, category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added expense: %s - $%.2f (%s)\n", title, amount, category)

	case "list":
		err := listExpenses()
		if err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: expense-tracker <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <title> <amount> <category>  Add a new expense")
	fmt.Println("  list                             List all expenses")
}

func parseAmount(s string) float64 {
	var amount float64
	_, err := fmt.Sscanf(s, "%f", &amount)
	if err != nil || amount <= 0 {
		log.Fatal("Invalid amount. Must be a positive number.")
	}
	return amount
}
