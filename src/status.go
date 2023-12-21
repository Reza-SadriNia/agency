package src

import (
	"agency/db"
	"agency/model"
	"fmt"
	"strconv"
	"strings"
)

func Status(region string) {
	data, err := db.LoadData()
	if err != nil {
		fmt.Println("Error loading data: ", err)

		return
	}

	var regionAgencues []model.Agency

	for _, item := range data {
		if item.Region == region {
			regionAgencues = append(regionAgencues, item)
		}
	}

	totalAgencies := len(regionAgencues)
	totalEmployee := 0
	for _, item := range regionAgencues {
		employeeCount, _ := strconv.Atoi(item.EmployeeCount)
		totalEmployee += employeeCount
	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(strings.Repeat("\n", 1))
	fmt.Printf("Total Agencies in %s: %d\n", region, totalAgencies)
	fmt.Printf("Total Employee in %s: %d\n", region, totalEmployee)
	fmt.Println(strings.Repeat("\n", 1))
	fmt.Println(strings.Repeat("-", 40))
}
