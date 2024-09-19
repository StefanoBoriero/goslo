package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello, Go!")
	http.HandleFunc("POST /burn", handler)
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", nil)
}

func calculate(sloWindow float64, errorBudgetPercentage float64, consumedInTime float64) float64 {
	return ((sloWindow * 24) / consumedInTime) * (errorBudgetPercentage / 100)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	window, _ := strconv.ParseFloat(r.Form.Get("slo-window"), 2)
	percentage, _ := strconv.ParseFloat(r.Form.Get("error-budget-percentage"), 2)
	time, _ := strconv.ParseFloat(r.Form.Get("time-consumption"), 2)
	fmt.Fprintf(w, strconv.FormatFloat(calculate(window, percentage, time), 'f', 2, 64))
}
