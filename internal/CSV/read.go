package CSV

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

func Read() ([][]string, error) {
	file, err := os.Open("C:/Users/Alireza/Desktop/business-financial-data-mar-2022-quarter-csv.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records, e := csv.NewReader(file).ReadAll()
	if e != nil {
		return nil, errors.New("could not read CSV file")
	}

	return records, nil

}
