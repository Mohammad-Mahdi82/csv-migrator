package CSV

import (
	"encoding/csv"
	"errors"
	"os"
)

func Read() ([][]string, error) {
	file, err := os.Open("C:/Users/Alireza/Desktop/business-financial-data-mar-2022-quarter-csv.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, e := csv.NewReader(file).ReadAll()
	if e != nil {
		return nil, errors.New("could not read CSV file")
	}

	return records, nil

}
