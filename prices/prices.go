package prices

import (
	"bufio"
	"fmt"
	"github.com/inuoluwadunsimi/price-calculator/conversion"
	"os"
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
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("reading file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("converting price to float failed")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
	file.Close()
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{20, 30, 40},
		Taxrate:     taxRate,
	}
}
