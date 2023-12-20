package db

import (
	"agency/model"
	"encoding/csv"
	"os"
	"strconv"
)

func SaveData(data []model.Agency) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		record := []string{
			strconv.Itoa(item.ID),
			item.Name,
			item.Region,
			item.Address,
			item.Phone,
			item.JoinDate,
			item.EmployeeCount,
		}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	return nil
}
