package main

import (
	"agency/src"
	"fmt"
	"strings"
)

func main() {
	for {
		fmt.Println("\nCommand:")
		fmt.Println("1. list - List agencies in a region")
		fmt.Println("2. get - Get details of an agency by ID")
		fmt.Println("3. create - Create a new agency")
		fmt.Println("4. edit - Edit details of an agency by ID")
		fmt.Println("5. status - Get status of agencies in a region")
		fmt.Println("6. exit - Exit the program")

		fmt.Print("Enter Command: ")
		command := src.GetUserInput()

		switch command {
		case "exit":
			return
		case "list":
		case "get":
		case "create":
		case "edit":
		case "status":
		default:
			fmt.Println(strings.Repeat("-", 30))
			fmt.Print(strings.Repeat("\n", 1))
			fmt.Println("Invalid command. Please try again.")
			fmt.Print(strings.Repeat("\n", 1))
			fmt.Println(strings.Repeat("-", 30))
		}
	}
}
