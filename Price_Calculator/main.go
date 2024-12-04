package main

import (
	"fmt"
	//"example.com/main/cmdmanager"
	"example.com/main/filemanager"
	"example.com/main/prices"
)

type IOManager interface{
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		// err := priceJob.Process()
		// if err != nil{
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
		go priceJob.Process(doneChans[index], errorChans[index])
	}
	for index, _ := range taxRates{
		select {
		case err := <- errorChans[index]:
			if err != nil{
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

	for _, errorChan := range errorChans{
		<- errorChan
	}
	for _, doneChan := range doneChans{
		<- doneChan
	}
}