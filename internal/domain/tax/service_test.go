package tax_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang-capital-gain/internal/domain/tax"
	"golang-capital-gain/internal/pkg/models"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name       string
		operations []models.Operation
		expected   []models.TaxOutput
	}{
		{
			name: "Buy operation",
			operations: []models.Operation{
				{Operation: models.BUY, Quantity: 10, UnitCost: 100},
			},
			expected: []models.TaxOutput{
				{Tax: 0},
			},
		},
		{
			name: "Sell operation with no tax",
			operations: []models.Operation{
				{Operation: models.BUY, Quantity: 10, UnitCost: 100},
				{Operation: models.SELL, Quantity: 10, UnitCost: 150},
			},
			expected: []models.TaxOutput{
				{Tax: 0},
				{Tax: 0},
			},
		},
		{
			name: "Sell operation with tax",
			operations: []models.Operation{
				{Operation: models.BUY, Quantity: 100, UnitCost: 100},
				{Operation: models.SELL, Quantity: 100, UnitCost: 300},
			},
			expected: []models.TaxOutput{
				{Tax: 0},
				{Tax: 4000}, // Gain: 20000, Tax: 20% of 20000
			},
		},
		{
			name: "Sell operation with loss",
			operations: []models.Operation{
				{Operation: models.BUY, Quantity: 100, UnitCost: 100},
				{Operation: models.SELL, Quantity: 100, UnitCost: 50},
				{Operation: models.SELL, Quantity: 100, UnitCost: 300},
			},
			expected: []models.TaxOutput{
				{Tax: 0},
				{Tax: 0},    // Loss: 5000
				{Tax: 3000}, // Gain: 20000 - 5000 = 15000, Tax: 20% of 15000
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taxService := tax.NewTaxService()

			result := taxService.CalculateTax(tt.operations)

			assert.Equal(t, tt.expected, result)
		})
	}
}
