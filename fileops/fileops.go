package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) int {
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func GetFloatFromFile(fileN string) float64 {

	data, fileError := os.ReadFile(fileN)
	b := handleError(fileError)
	if b == 1 {
		return 1000
	}

	valueText := string(data)
	value, floatError := strconv.ParseFloat(valueText, 64)
	handleError(floatError)

	return value
}

func WriteFloatToFile(value float64, fileN string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileN, []byte(valueText), 0644)
}
