package core

import (
	"CSV/pkg/db"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func InsertDataChunk(base int8, data *[][]string, loader ILoader) {

	var i int

	for i = 0; i < (len(*data) / int(base)); i++ {
		fmt.Printf("pack cursor: %v\n", i*10)
		loader(base, data, i*int(base), InsertRow)
	}

	if diff := len(*data) % int(base); diff != 0 {
		fmt.Printf("triggered! diff:%v\n", diff)
		loader(int8(diff), data, i*int(base), InsertRow)
	}

}

func InsertManyRows(base int8, data *[][]string, cur int, loader IWorkLoader) {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < int(base); i++ {
			wg.Add(1)
			go func(data *[][]string, in int) {
				defer wg.Done()
				result := loader(data, in)
				dataChan <- result
			}(data, cur+i)
		}
		wg.Wait()
		close(dataChan)
	}()

	<-dataChan

}

func InsertRow(records *[][]string, i int) int {
	fmt.Printf("index: %v\n", i+1)

	data, _ := Convert(records, i)
	_, erro := db.FinancialDataCollection.InsertOne(context.Background(), data)
	if erro != nil {
		log.Fatalf("Can not write to mongo: %v\n", erro)
	}

	time.Sleep(50 * time.Millisecond)
	return i
}
