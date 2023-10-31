package core

import "CSV/pkg/db"

func Convert(records *[][]string, i int) (db.FinancialDataType, error) {
	Records := *records
	record := Records[i]

	data := db.FinancialDataType{
		SeriesReference: record[0],
		Period:          record[1],
		DataValue:       record[2],
		Suppressed:      record[3],
		Status:          record[4],
		Units:           record[5],
		Magnitude:       record[6],
		Subject:         record[7],
		Group:           record[8],
		SeriesTitle1:    record[9],
		SeriesTitle2:    record[10],
		SeriesTitle3:    record[11],
		SeriesTitle4:    record[12],
		SeriesTitle5:    record[13],
	}
	return data, nil
}
