package prices

import (
	"fmt"
	"example.com/main/conversion"
	"example.com/main/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices`
	TaxIncludedPrices map[string]float64 `json:"tax_included_price`
}

func NewTaxIncludedPriceJob(taxRate float64, iom iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	//errorChan <- errors.New("An error")

	if err != nil{
		//return err
		errorChan <- err
		return
	}

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	
	fmt.Println(result)
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error{
	lines, err := job.IOManager.ReadLines()
	
	if err != nil{
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil{
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

