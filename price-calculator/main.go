package main

import (
	"fmt"

	"github.com/master-bogdan/price-calculator/filemanager"
	"github.com/master-bogdan/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChannels := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)

		var inputPath string = "prices.txt"
		var outputPath string = fmt.Sprintf("result_%.0f.json", taxRate*100)

		fm := filemanager.New(inputPath, outputPath)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[index])
	}

	for _, doneChannel := range doneChannels {
		<-doneChannel
	}
}
