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
			fillemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}

	fmt.Println(result)

}
