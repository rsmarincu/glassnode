package repository

import (
	"context"
	"database/sql"
)

type ETHRepository struct {
	db *sql.DB
}

func NewETHRepository(db *sql.DB) *ETHRepository {
	return &ETHRepository{
		db: db,
	}
}

func (r *ETHRepository) QueryEOATransactions(ctx context.Context) ([]*Transaction, error) {
	query := `select * from transactions t where not exists (
	select address from contracts c 
	where t.from = c.address and t.to = c.address and c.address = '0x0000000000000000000000000000000000000000'
)`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*Transaction, 0)
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction)
		if err != nil {
			return nil, err
		}
		result = append(result, &transaction)
	}

	return result, nil
}
