package tax

import (
	"math"

	"golang-capital-gain/internal/domain/stocks"
	"golang-capital-gain/internal/pkg/models"
	"golang-capital-gain/internal/utils"
)

type TaxService struct {
	Losses float64
}

func NewTaxService(stocksService *stocks.StocksService) *TaxService {
	return &TaxService{
		Losses: 0,
	}
}

func (s *TaxService) CalculateTax(operations []models.Operation) []models.TaxOutput {
	var tax []models.TaxOutput

	stocksService := stocks.NewStocksService()

	for _, operation := range operations {
		switch operation.Operation {
		case models.BUY:
			stocksService.Buy(operation)
			tax = append(tax, buildTax(0))

		case models.SELL:
			taxValue := 0.0
			gain, loss := stocksService.Sell(operation)
			gain = s.deductLoss(gain, loss)

			if gain > 0 {
				taxValue = getTaxValue(gain, operation)
			}

			tax = append(tax, buildTax(taxValue))
		}
	}

	return tax
}

func (s *TaxService) deductLoss(gain, loss float64) float64 {
	currLoss := s.Losses

	if loss > 0 {
		s.Losses += loss
	}

	if gain > 0 {
		s.Losses -= gain

		if s.Losses < 0 {
			s.Losses = 0
		}
	}

	return gain - currLoss
}

func buildTax(value float64) models.TaxOutput {
	return models.TaxOutput{
		Tax: value,
	}
}

func getTaxValue(currGain float64, operation models.Operation) float64 {
	if float64(operation.Quantity)*(operation.UnitCost) <= 20000 {
		return 0
	}

	return math.Abs(utils.ApplyPrecision(currGain*0.2, 2))
}
