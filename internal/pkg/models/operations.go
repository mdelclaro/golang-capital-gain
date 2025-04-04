package models

type (
	Operation struct {
		Operation OperationType `json:"operation"`
		UnitCost  float64       `json:"unit-cost"`
		Quantity  int           `json:"quantity"`
	}

	OperationType string
)

const (
	BUY  OperationType = "buy"
	SELL OperationType = "sell"
)
