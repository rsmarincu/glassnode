package usecases_test

import (
	"context"
	"errors"
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

		var limit uint32 = 100
		var offset uint32 = 0

		mockRepository := mock_usecases.NewMockETHRepository(ctrl)
		feesService := usecases.NewFeesService(mockRepository)

		transactions := getTestTransactions(100)
		mockRepository.EXPECT().QueryEOATransactions(gomock.Any()).Return(transactions, nil)

		fees, moreResults, err := feesService.ListFees(ctx, offset, limit)

		require.NoError(t, err)
		assert.NotEmpty(t, fees)
		assert.False(t, moreResults)
		assert.Equal(t, 0.0, fees[0].Value)
	})

	t.Run("successfully offsets", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var limit uint32 = 100
		var offset uint32 = 1

		mockRepository := mock_usecases.NewMockETHRepository(ctrl)
		feesService := usecases.NewFeesService(mockRepository)

		transactions := getTestTransactions(100)
		mockRepository.EXPECT().QueryEOATransactions(gomock.Any()).Return(transactions, nil)

		fees, moreResults, err := feesService.ListFees(ctx, offset, limit)

		require.NoError(t, err)
		assert.NotEmpty(t, fees)
		assert.False(t, moreResults)
		assert.Equal(t, 0.0, fees[0].Value)
		assert.Len(t, fees, 1)

	})

	t.Run("propagates error from storage", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var limit uint32 = 100
		var offset uint32 = 0

		mockRepository := mock_usecases.NewMockETHRepository(ctrl)
		feesService := usecases.NewFeesService(mockRepository)

		mockRepository.EXPECT().QueryEOATransactions(gomock.Any()).Return(nil, errors.New("database error"))

		fees, moreResults, err := feesService.ListFees(ctx, offset, limit)

		require.Error(t, err)
		assert.Empty(t, fees)
		assert.False(t, moreResults)
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
