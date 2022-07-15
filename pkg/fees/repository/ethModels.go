package repository

import "time"

type Transaction struct {
	TxID        string
	BlockHeight int64
	BlockHash   string
	BlockTime   time.Time
	From        string
	To          string
	Value       float64
	GasProvided float64
	GasUsed     float64
	GasPrice    float64
	Status      string
}
