package repository

import "time"

type Transaction struct {
	TxID        string    `db:"txid"`
	BlockHeight int64     `db:"block_height"`
	BlockHash   string    `db:"block_hash"`
	BlockTime   time.Time `db:"block_time"`
	From        string    `db:"from"`
	To          string    `db:"to"`
	Value       float64   `db:"value"`
	GasProvided float64   `db:"gas_provided"`
	GasUsed     float64   `db:"gas_used"`
	GasPrice    float64   `db:"gas_price"`
	Status      string    `db:"status"`
}
