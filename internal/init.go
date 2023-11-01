package internal

import (
	"CSV/internal/CSV"
	"CSV/internal/core"
	"log"
)

//

func Migrate() {
	records, err := CSV.Read()
	if err != nil {
		log.Fatal(err)
	}

	core.InsertDataChunk(10, &records, core.InsertManyRows)
}
