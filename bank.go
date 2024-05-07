package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

type Bank struct {
	balance float64
}

func initializeBank(params ...float64) *Bank {
	var balance float64 = 0
	if params[0] > 0 {
		balance = params[0]
	}
	return &Bank{balance: balance}
}
func main() {
	balance := fileops.GetFloatFromFile(fileName)
	bank := initializeBank(balance)
	fmt.Printf("Reach us 24/7: %s\n", randomdata.PhoneNumber())
	handleBankOperations(bank)
}

func checkBalance(bank *Bank) {
	fmt.Printf("Your balance is: %.2f\n", bank.balance)
}

func inputAndValidateAmount(text string) float64 {
	var amount float64 = 0
	for amount <= 0 {
		fmt.Print(text)
		fmt.Scan(&amount)
	}
	return amount
}
func depositMoney(bank *Bank) {
	bank.balance += inputAndValidateAmount("Deposit a valid amount: ")
	fileops.WriteFloatToFile(bank.balance, fileName)
}

func withdrawMoney(bank *Bank) {
	bank.balance -= inputAndValidateAmount("Withdraw a valid amount: ")
	fileops.WriteFloatToFile(bank.balance, fileName)
}

func handleBankOperations(bank *Bank) {
	var userInput uint = 0
	for {
		displayMenu(&userInput)
		switch userInput {
		case 1:
			checkBalance(bank)
		case 2:
			depositMoney(bank)
		case 3:
			withdrawMoney(bank)
		case 4:
			fmt.Println("You chose to exit")
			return
		default:
			fmt.Println("Please enter a valid operation")
		}
	}
}
