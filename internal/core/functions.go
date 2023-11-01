package core

import (
	"CSV/pkg/db"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

//

func InsertDataChunk(base int8, data *[][]string, loader ILoader) {

	var i int

	for i = 0; i < (len(*data) / int(base)); i++ {
		fmt.Printf("pack cursor: %v\n", i*10)
		err := loader(base, data, i*int(base), InsertRow)
		if err != nil {
			log.Fatal(err)
		}
	}

	if diff := len(*data) % int(base); diff != 0 {
		fmt.Printf("triggered! diff:%v\n", diff)
		err := loader(int8(diff), data, i*int(base), InsertRow)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func InsertManyRows(base int8, data *[][]string, cur int, loader IWorkLoader) error {
	dataChan := make(chan int)
	var e error

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < int(base); i++ {
			wg.Add(1)
			go func(data *[][]string, in int) {

				err := func() error {
					defer wg.Done()
					result, err := loader(data, in)
					if err != nil {
						return err
					}
					dataChan <- result
					return nil
				}()

				if err != nil {
					e = err
				}

			}(data, cur+i)
		}
		wg.Wait()
		close(dataChan)
	}()

	<-dataChan

	if e != nil {
		return e
	}

	return nil

}

func InsertRow(records *[][]string, i int) (int, error) {
	fmt.Printf("index: %v\n", i+1)

	data, _ := Convert(records, i)
	_, erro := db.FinancialDataCollection.InsertOne(context.Background(), data)
	if erro != nil {
		return -1, erro
	}

	time.Sleep(50 * time.Millisecond)
	return i, nil
}
