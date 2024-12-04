package main

import (
	"fmt"
	"example.com/bank/fileops"
	"example.com/bank/communication"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		communication.Option()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += deposit
			fmt.Println("Balance updated! New account: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdraw float64
			fmt.Print("Withdraw: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdraw
			fmt.Println("Balance updated! New account: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	var deposit float64
		// 	fmt.Print("Your deposit: ")
		// 	fmt.Scan(&deposit)

		// 	if deposit <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance += deposit
		// 	fmt.Println("Balance updated! New account: ", accountBalance)
		// } else if choice == 3 {
		// 	var withdraw float64
		// 	fmt.Print("Withdraw: ")
		// 	fmt.Scan(&withdraw)

		// 	if withdraw <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdraw > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdraw
		// 	fmt.Println("Balance updated! New account: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	fmt.Println("Thanks for choosing our bank")
		// 	break
		// }
	}

}
