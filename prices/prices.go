package prices

import (
	"fmt"
	"github.com/inuoluwadunsimi/price-calculator/conversion"
	"github.com/inuoluwadunsimi/price-calculator/fillemanager"
)

type TaxIncludedPriceJob struct {
	Taxrate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.Taxrate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}
	fmt.Println(result)

}

func (job *TaxIncludedPriceJob) loadData() {

	lines, err := fillemanager.Readlines("prices.txt")
	if err != nil {
		fmt.Println(err)
	}
	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("converting price to float failed")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{20, 30, 40},
		Taxrate:     taxRate,
	}
}
