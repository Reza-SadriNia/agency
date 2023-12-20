package db

import (
	"fmt"
	"os"
)

const dataFile = "./db/agencies.txt"

func init() {
	var file *os.File

	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		fmt.Println("can`t create or open file", err)

		return
	}
	defer file.Close()
}
