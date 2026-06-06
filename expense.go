package main

type Expense struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Note     string  `json:"note"`
	Category string  `json:"category"`
	Date     string  `json:"created_at"`
}
