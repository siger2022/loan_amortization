package main

import (
	"fmt"
)

func main() {

	var initial_loan, loan, payment, interest float32

	for loan >= 0 {
		fmt.Println("Type in a loan: ")
		fmt.Println("Or enter a negative number if you want to stop running the program")
		fmt.Scanln(&loan)

		switch {

		case loan < 0:
			fmt.Println("The program is terminated")
			break

		case loan >= 0:
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
			fmt.Printf("For a loan of %v with the annual interest rate of %v\n", initial_loan, interest)
			fmt.Printf("With an average monthy payment of: %v ", payment)

			if month < 12 {
				fmt.Println("It will require: ", month, "months")
			} else {
				year := month / 12
				months := month % 12
				fmt.Printf("It will require: %v years and %v month or %v months\n", year, months, month)
			}
		}
	}
}
