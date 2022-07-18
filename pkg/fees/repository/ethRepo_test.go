package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rsmarincu/glassnode/pkg/fees/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestQuery(t *testing.T) {
	t.Run("successfully returns rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error occured while opening a mock database connection: %s", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"txid", "block_height", "block_hash", "block_time", "from", "to", "value", "gas_provided", "gas_used", "gas_price", "status"}).
			AddRow("id_1", 1, "hash_1", time.Now(), "from_1", "to_1", 1, 1, 1, 1, "true").
			AddRow("id_2", 2, "hash_2", time.Now(), "from_2", "to_2", 1, 1, 1, 1, "true").
			AddRow("id_3", 3, "hash_3", time.Now(), "from_3", "to_3", 1, 1, 1, 1, "true")

		query := `select * from transactions t where not exists (
	select address from contracts c 
	where t.from = c.address and t.to = c.address and c.address = '0x0000000000000000000000000000000000000000'
) order by t.block_time`

		mock.ExpectQuery(query).WillReturnRows(rows)

		ethRepo := repository.NewETHRepository(db)
		transactions, err := ethRepo.QueryEOATransactions(context.Background())

		assert.NoError(t, err)
		assert.Len(t, transactions, 3)
	})

}
