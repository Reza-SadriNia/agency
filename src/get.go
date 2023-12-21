package src

import (
	"agency/db"
	"agency/model"
	"fmt"
	"strings"
)

func GetAgencyDetail(agencyID int) {
	data, err := db.LoadData()
	if err != nil {
		fmt.Println("Error loading data: ", err)
		return
	}

	var foundAgency model.Agency

	for _, item := range data {
		if item.ID == agencyID {
			foundAgency = item
			break
		}
	}

	if foundAgency.ID != 0 {
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println(strings.Repeat("\n", 1))
		fmt.Println("Agency Details:")
		fmt.Printf("ID: %d, Name: %s, Address: %s, Phone: %s, Join Date: %s, Employee Count: %s\n",
			foundAgency.ID,
			foundAgency.Name,
			foundAgency.Address,
			foundAgency.Phone,
			foundAgency.JoinDate,
			foundAgency.EmployeeCount,
		)
		fmt.Println(strings.Repeat("\n", 1))
		fmt.Println(strings.Repeat("-", 40))
	} else {
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println(strings.Repeat("\n", 1))
		fmt.Println("Agency not Found.")
		fmt.Println(strings.Repeat("\n", 1))
		fmt.Println(strings.Repeat("-", 40))
	}

}
