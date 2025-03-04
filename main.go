package main

import (
	"fmt"
	"github.com/inuoluwadunsimi/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(result)

}
