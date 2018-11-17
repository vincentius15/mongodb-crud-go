package main

type exchange struct {
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}
