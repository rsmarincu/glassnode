package grpc_test

import (
	"context"
	"errors"
	feespb "github.com/rsmarincu/glassnode/api"
	"github.com/rsmarincu/glassnode/pkg/fees"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/rsmarincu/glassnode/pkg/fees/grpc"
	mock_grpc "github.com/rsmarincu/glassnode/pkg/fees/grpc/mocks"
)

func TestFees(t *testing.T) {
	t.Run("successfully returns fees", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		mockFeesService := mock_grpc.NewMockFeesService(ctrl)

		handler := grpc.NewServiceHandler(mockFeesService)

		req := feespb.ListFeesRequest{
			PageSize:  0,
			PageToken: "",
		}
		var offset, limit uint32 = 0, 10
		fees := []*fees.Fee{
			{
				Timestamp: time.Now(),
				Value:     1,
			},
			{
				Timestamp: time.Now(),
				Value:     2,
			},
			{
				Timestamp: time.Now(),
				Value:     3,
			},
			{
				Timestamp: time.Now(),
				Value:     4,
			},
			{
				Timestamp: time.Now(),
				Value:     5,
			},
		}

		mockFeesService.EXPECT().ListFees(gomock.Any(), offset, limit).Return(fees, false, nil)
		response, err := handler.ListFees(ctx, &req)

		require.NoError(t, err)
		assert.Len(t, response.Fees, len(fees))
	})

	t.Run("successfully returns response with no fees", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		mockFeesService := mock_grpc.NewMockFeesService(ctrl)

		handler := grpc.NewServiceHandler(mockFeesService)

		req := feespb.ListFeesRequest{
			PageSize:  0,
			PageToken: "",
		}
		var offset, limit uint32 = 0, 10

		mockFeesService.EXPECT().ListFees(gomock.Any(), offset, limit).Return(nil, false, nil)
		response, err := handler.ListFees(ctx, &req)

		require.NoError(t, err)
		assert.Empty(t, response.Fees)
	})

	t.Run("propagates service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		mockFeesService := mock_grpc.NewMockFeesService(ctrl)

		handler := grpc.NewServiceHandler(mockFeesService)

		req := feespb.ListFeesRequest{
			PageSize:  0,
			PageToken: "",
		}
		var offset, limit uint32 = 0, 10

		mockFeesService.EXPECT().ListFees(gomock.Any(), offset, limit).Return(nil, false, errors.New("new service error"))
		response, err := handler.ListFees(ctx, &req)

		require.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("successfully shows next page token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		mockFeesService := mock_grpc.NewMockFeesService(ctrl)

		handler := grpc.NewServiceHandler(mockFeesService)

		req := feespb.ListFeesRequest{
			PageSize:  2,
			PageToken: "",
		}
		var offset, limit uint32 = 0, 2
		fees := []*fees.Fee{
			{
				Timestamp: time.Now(),
				Value:     1,
			},
			{
				Timestamp: time.Now(),
				Value:     2,
			},
		}

		mockFeesService.EXPECT().ListFees(gomock.Any(), offset, limit).Return(fees, true, nil)
		response, err := handler.ListFees(ctx, &req)

		require.NoError(t, err)
		assert.NotEmpty(t, response.NextPageToken)
		assert.Len(t, response.Fees, len(fees))
	})
}
