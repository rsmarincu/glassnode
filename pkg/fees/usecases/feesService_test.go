package usecases_test

import (
	"context"
	"fmt"
	"github.com/rsmarincu/glassnode/pkg/fees/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/rsmarincu/glassnode/pkg/fees/usecases"
	mock_usecases "github.com/rsmarincu/glassnode/pkg/fees/usecases/mocks"
)

func TestFeesService_ListFees(t *testing.T) {
	t.Run("successfully computes values", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepository := mock_usecases.NewMockETHRepository(ctrl)
		feesService := usecases.NewFeesService(mockRepository)

		transactions := getTestTransactions(100)
		mockRepository.EXPECT().QueryEOATransactions(gomock.Any()).Return(transactions, nil)

		fees, err := feesService.ListFees(ctx)

		require.NoError(t, err)
		assert.NotEmpty(t, fees)
		assert.Equal(t, 0.0, fees[0].Value)
	})
}

func getTestTransactions(length int) []*repository.Transaction {
	transactions := make([]*repository.Transaction, length)
	now := time.Now()

	for i := 0; i < length; i++ {
		transactions[i] = &repository.Transaction{
			TxID:      fmt.Sprintf("txid-%d", i),
			GasPrice:  0,
			GasUsed:   1,
			BlockTime: now.Add(time.Duration(i) * time.Minute),
		}
	}
	return transactions
}
