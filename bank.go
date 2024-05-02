package main

import "fmt"

type Bank struct {
	balance float64
}

func initializeBank() *Bank {
	return &Bank{balance: 0.0}
}
func main() {
	bank := initializeBank()
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
}

func withdrawMoney(bank *Bank) {
	bank.balance -= inputAndValidateAmount("Withdraw a valid amount: ")
}

func displayMenu(userChoice *uint) {
	fmt.Println("Welcome to GO bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Please, input your choice: ")
	fmt.Scan(userChoice)
	fmt.Print("-------------------------------\n")
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
