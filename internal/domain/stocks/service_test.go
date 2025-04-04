package stocks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang-capital-gain/internal/domain/stocks"
	"golang-capital-gain/internal/pkg/models"
)

func TestStocksService(t *testing.T) {
	t.Run("Buy operations", func(t *testing.T) {
		stocksService := stocks.NewStocksService()

		// Perform a buy operation
		stocksService.Buy(models.Operation{
			Operation: models.BUY,
			Quantity:  10,
			UnitCost:  100.00,
		})

		assert.Equal(t, 10, stocksService.Shares, "Shares count mismatch after buy")
		assert.Equal(t, 100.00, stocksService.AveragePrice, "Average price mismatch after buy")

		// Perform another buy operation
		stocksService.Buy(models.Operation{
			Operation: models.BUY,
			Quantity:  20,
			UnitCost:  200.00,
		})

		assert.Equal(t, 30, stocksService.Shares, "Shares count mismatch after second buy")
		assert.Equal(t, 166.67, stocksService.AveragePrice, "Average price mismatch after second buy")
	})

	t.Run("Sell operations", func(t *testing.T) {
		stocksService := stocks.NewStocksService()

		// Perform buy operations
		stocksService.Buy(models.Operation{
			Operation: models.BUY,
			Quantity:  10,
			UnitCost:  100.00,
		})
		stocksService.Buy(models.Operation{
			Operation: models.BUY,
			Quantity:  20,
			UnitCost:  200.00,
		})

		// Perform a sell operation with gain
		gain, loss := stocksService.Sell(models.Operation{
			Operation: models.SELL,
			Quantity:  10,
			UnitCost:  300.00,
		})

		assert.Equal(t, 20, stocksService.Shares, "Shares count mismatch after sell")
		assert.Equal(t, 1333.3, gain, "Gain mismatch after sell")
		assert.Equal(t, 0.00, loss, "Loss mismatch after sell")

		// Perform a sell operation with loss
		gain, loss = stocksService.Sell(models.Operation{
			Operation: models.SELL,
			Quantity:  10,
			UnitCost:  50.00,
		})

		assert.Equal(t, 10, stocksService.Shares, "Shares count mismatch after second sell")
		assert.Equal(t, 0.00, gain, "Gain mismatch after second sell")
		assert.Equal(t, 1166.7, loss, "Loss mismatch after second sell")

		// Perform a sell operation with no gain or loss
		gain, loss = stocksService.Sell(models.Operation{
			Operation: models.SELL,
			Quantity:  10,
			UnitCost:  166.67,
		})

		assert.Equal(t, 0, stocksService.Shares, "Shares count mismatch after third sell")
		assert.Equal(t, 0.00, gain, "Gain mismatch after third sell")
		assert.Equal(t, 0.00, loss, "Loss mismatch after third sell")
	})
}
