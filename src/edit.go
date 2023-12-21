package src

import (
	"agency/db"
	"agency/model"
	"fmt"
	"strings"
)

func EditAgency(agencyID int) {
	data, err := db.LoadData()
	if err != nil {
		fmt.Println("Error loading data: ", err)

		return
	}

	var foundAgency model.Agency

	for i, item := range data {
		if item.ID == agencyID {
			foundAgency = item
			foundAgencyIndex := i

			fmt.Printf("Editing Agency %d:\n", agencyID)
			fmt.Printf("Name (%s): ", foundAgency.Name)
			newName := GetUserInput()
			if newName != "" {
				foundAgency.Name = newName
			}

			fmt.Printf("Region (%s): ", foundAgency.Region)
			newRegion := GetUserInput()
			if newRegion != "" {
				foundAgency.Region = newRegion
			}

			fmt.Printf("Address (%s): ", foundAgency.Address)
			newAddress := GetUserInput()
			if newAddress != "" {
				foundAgency.Address = newAddress
			}

			fmt.Printf("Phone (%s): ", foundAgency.Phone)
			newPhone := GetUserInput()
			if newPhone != "" {
				foundAgency.Phone = newPhone
			}

			fmt.Printf("EmployeeCount (%s): ", foundAgency.EmployeeCount)
			newEmployeeCount := GetUserInput()
			if newEmployeeCount != "" {
				foundAgency.EmployeeCount = newEmployeeCount
			}

			data[foundAgencyIndex] = foundAgency

			err = db.SaveData(data)
			if err != nil {
				fmt.Println("Error saving data:", err)

				return
			}

			fmt.Println(strings.Repeat("-", 40))
			fmt.Println(strings.Repeat("\n", 1))
			fmt.Println("Agency edited successfully.")
			fmt.Println(strings.Repeat("\n", 1))
			fmt.Println(strings.Repeat("-", 40))

			return
		}
	}
	fmt.Println("Agency not found.")
}
