package internal

import (
	"CSV/internal/core"
	"fmt"
	"log"
	"testing"
	"time"
)

var sample [][]string = [][]string{

	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
	{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
	},
}

func sampleRowLoader(records *[][]string, i int) (int, error) {
	fmt.Printf("index: %v\n", i+1)
	time.Sleep(50 * time.Millisecond)
	return i, nil
}

func TestConvert(t *testing.T) {
	t.Parallel()
	_, err := core.Convert(&sample, 0)
	if err != nil {
		t.Fail()
	}
}

func TestInsertData(t *testing.T) {

	var base int8 = 10

	var i int

	for i = 0; i < (len(sample) / int(base)); i++ {
		fmt.Printf("pack cursor: %v\n", i*10)
		err := core.InsertManyRows(base, &sample, i*int(base), sampleRowLoader)
		if err != nil {
			log.Fatal(err)
		}
	}

	if diff := len(sample) % int(base); diff != 0 {
		fmt.Printf("triggered! diff:%v\n", diff)
		err := core.InsertManyRows(int8(diff), &sample, i*int(base), sampleRowLoader)
		if err != nil {
			log.Fatal(err)
		}
	}

}
