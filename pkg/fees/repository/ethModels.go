package repository

import "time"

type Transaction struct {
	TxID        string
	BlockHeight int8
	BlockHash   string
	BlockTime   time.Time
	From        string
	To          string
	Value       float32
	GasProvided float32
	GasUsed     float32
	Status      string
}
