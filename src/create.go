package src

import (
	"agency/db"
	"agency/model"
	"fmt"
	"strings"
	"time"
)

func CreateAgencies(name, region, address, phone, employeeCount string) {
	data, err := db.LoadData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	currentTime := time.Now()
	newAgency := model.Agency{
		ID:            len(data) + 1,
		Name:          name,
		Region:        region,
		Address:       address,
		Phone:         phone,
		JoinDate:      currentTime.Format("2006/01/02 15:04:05"),
		EmployeeCount: employeeCount,
	}
	data = append(data, newAgency)

	err = db.SaveData(data)

	if err != nil {
		fmt.Println("error saving data:", err)
		return
	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(strings.Repeat("\n", 1))
	fmt.Println("Agency created successfully")
	fmt.Println(strings.Repeat("\n", 1))
	fmt.Println(strings.Repeat("-", 40))
}
