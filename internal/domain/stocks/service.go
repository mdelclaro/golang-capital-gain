package stocks

import (
	"math"

	"golang-capital-gain/internal/pkg/models"
	"golang-capital-gain/internal/utils"
)

type StocksService struct {
	AveragePrice float64
	Shares       int
}

func NewStocksService() *StocksService {
	return &StocksService{
		AveragePrice: utils.ZERO_VALUE,
		Shares:       utils.ZERO_VALUE,
	}
}

func (s *StocksService) Buy(stock models.Operation) {
	s.AveragePrice = ((float64(s.Shares) * s.AveragePrice) +
		(float64(stock.Quantity) * stock.UnitCost)) /
		float64(s.Shares+stock.Quantity)

	s.Shares += stock.Quantity

	s.AveragePrice = utils.ApplyPrecision(s.AveragePrice, 2)
}

// returns gain and loss respectively
func (s *StocksService) Sell(operation models.Operation) (float64, float64) {
	operationCost := float64(operation.Quantity) * operation.UnitCost
	currentCost := float64(operation.Quantity) * s.AveragePrice

	s.Shares -= operation.Quantity

	diff := utils.ApplyPrecision(currentCost-operationCost, 2)

	if operation.UnitCost == s.AveragePrice {
		return utils.ZERO_VALUE, utils.ZERO_VALUE
	}

	if operation.UnitCost < s.AveragePrice {
		return utils.ZERO_VALUE, diff
	}

	return math.Abs(diff), utils.ZERO_VALUE
}
