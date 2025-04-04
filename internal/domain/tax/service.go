package tax

import (
	"math"

	"golang-capital-gain/internal/domain/stocks"
	"golang-capital-gain/internal/pkg/models"
	"golang-capital-gain/internal/utils"
)

const TAX_PERCENT = float64(0.2)

type TaxService struct {
	Losses float64
}

func NewTaxService() *TaxService {
	return &TaxService{
		Losses: utils.ZERO_VALUE,
	}
}

func (s *TaxService) CalculateTax(operations []models.Operation) []models.TaxOutput {
	var taxes []models.TaxOutput

	stocksService := stocks.NewStocksService()

	for _, operation := range operations {
		switch operation.Operation {
		case models.BUY:
			stocksService.Buy(operation)
			taxes = append(taxes, buildTax(utils.ZERO_VALUE))

		case models.SELL:
			taxValue := utils.ZERO_VALUE
			gain, loss := stocksService.Sell(operation)
			gain = s.deductLoss(gain, loss)

			if gain > utils.ZERO_VALUE {
				taxValue = getTaxValue(gain, operation)
			}

			taxes = append(taxes, buildTax(taxValue))
		}
	}

	return taxes
}

func (s *TaxService) deductLoss(gain, loss float64) float64 {
	currLoss := s.Losses

	if loss > utils.ZERO_VALUE {
		s.Losses += loss
	}

	if gain > utils.ZERO_VALUE {
		s.Losses -= gain

		if s.Losses < utils.ZERO_VALUE {
			s.Losses = utils.ZERO_VALUE
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
		return utils.ZERO_VALUE
	}

	return math.Abs(utils.ApplyPrecision(currGain*TAX_PERCENT, 2))
}
