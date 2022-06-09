package main

import "fmt"

func main() {
	var initial_loan, loan, payment, interest float32
	fmt.Print("Type in a loan: ")
	fmt.Scan(&loan)

	initial_loan = loan

	fmt.Print("Average monthly payment: ")
	fmt.Scan(&payment)
	fmt.Print("Input annual interest rate: ")
	fmt.Scan(&interest)
	month := 0
	for loan > 0 {
		loan = loan * (1 + (interest / (100 * 12)))
		loan = loan - payment
		fmt.Println(loan)
		month += 1
	}

	fmt.Printf("For a loan of %v with the annual interest rate of %v%\n", initial_loan, interest)
	fmt.Println("With an average monthy payment of: ", payment, ".It will require:", month, "months ")

}
