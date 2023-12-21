package src

import (
	"agency/db"
	"fmt"
	"strings"
)

func ListAgencies(region string) {
	data, err := db.LoadData()
	if err != nil {
		fmt.Println("Error loading data:", err)
	}

	for _, item := range data {
		if item.Region == region {
			fmt.Println(strings.Repeat("-", 40))
			fmt.Println(strings.Repeat("\n", 1))
			fmt.Printf("ID: %d, Name: %s, Address: %s", item.ID, item.Name, item.Address)
			fmt.Println(strings.Repeat("\n", 1))
			fmt.Println(strings.Repeat("-", 40))
		}
	}
}
