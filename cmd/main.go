package main

import (
	"bufio"
	"encoding/json"
	"os"

	"golang-capital-gain/internal/domain/tax"
	"golang-capital-gain/internal/pkg/models"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		taxService := tax.NewTaxService()

		data := []models.Operation{}

		json.Unmarshal([]byte(scanner.Text()), &data)

		taxes := taxService.CalculateTax(data)

		json.NewEncoder(os.Stdout).Encode(taxes)
	}
}
