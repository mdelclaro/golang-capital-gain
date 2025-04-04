package main

import (
	"bufio"
	"encoding/json"
	"os"

	"golang-capital-gain/internal/domain/stocks"
	"golang-capital-gain/internal/domain/tax"
	"golang-capital-gain/internal/pkg/models"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		stocksService := stocks.NewStocksService()
		taxService := tax.NewTaxService(stocksService)

		data := []models.Operation{}

		json.Unmarshal([]byte(scanner.Text()), &data)

		capitalGainOutputs := taxService.CalculateTax(data)

		json.NewEncoder(os.Stdout).Encode(capitalGainOutputs)
	}
}
