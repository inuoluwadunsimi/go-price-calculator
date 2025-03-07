package prices

import (
	"fmt"
	"github.com/inuoluwadunsimi/price-calculator/conversion"
	"github.com/inuoluwadunsimi/price-calculator/fillemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         fillemanager.FileManager
	Taxrate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.Taxrate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}
	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) loadData() {

	lines, err := job.IOManager.Readlines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("converting price to float failed")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(fm fillemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{20, 30, 40},
		Taxrate:     taxRate,
	}
}
