package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

type Bank struct {
	balance float64
}

func handleError(err error) int {
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func getBalancefromFile() float64 {

	data, fileError := os.ReadFile(fileName)
	b := handleError(fileError)
	if b == 1 {
		return 1000
	}

	balanceText := string(data)
	balance, floatError := strconv.ParseFloat(balanceText, 64)
	handleError(floatError)

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func initializeBank(params ...float64) *Bank {
	var balance float64 = 0
	if params[0] > 0 {
		balance = params[0]
	}
	return &Bank{balance: balance}
}
func main() {
	balance := getBalancefromFile()
	bank := initializeBank(balance)
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
	writeBalanceToFile(bank.balance)
}

func withdrawMoney(bank *Bank) {
	bank.balance -= inputAndValidateAmount("Withdraw a valid amount: ")
	writeBalanceToFile(bank.balance)
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
