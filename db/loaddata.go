package db

import (
	"agency/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadData() ([]model.Agency, error) {
	var data []model.Agency

	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fields := strings.Split(line, ",")
		id, _ := strconv.Atoi(fields[0])
		data = append(data, model.Agency{
			ID:            id,
			Name:          fields[1],
			Region:        fields[2],
			Address:       fields[3],
			Phone:         fields[4],
			JoinDate:      fields[5],
			EmployeeCount: fields[6],
		})
	}
	return data, nil
}
