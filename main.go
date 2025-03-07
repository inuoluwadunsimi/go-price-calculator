package main

import (
	"fmt"
	"github.com/inuoluwadunsimi/price-calculator/fillemanager"
	"github.com/inuoluwadunsimi/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm :=
			fillemanager.New("prices.txt", fmt.Sprintf("result_%0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

	fmt.Println(result)

}
