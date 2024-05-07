package main

import "fmt"

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
