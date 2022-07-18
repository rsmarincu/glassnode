package grpc

import (
	"context"
	"fmt"
	"github.com/rsmarincu/glassnode/pkg/common"

	feespb "github.com/rsmarincu/glassnode/api"
	"github.com/rsmarincu/glassnode/pkg/fees"
)

const defaultPageSize = 10

type FeesService interface {
	ListFees(ctx context.Context, offset uint32, limit uint32) ([]*fees.Fee, bool, error)
}

type ServiceHandler struct {
	feespb.UnimplementedFeesServer
	feesService FeesService
}

func NewServiceHandler(feesService FeesService) feespb.FeesServer {
	return &ServiceHandler{
		feesService: feesService,
	}
}

func (s *ServiceHandler) ListFees(ctx context.Context, req *feespb.ListFeesRequest) (*feespb.ListFeesResponse, error) {
	offset, err := common.DecodeToken(req.PageToken)
	if err != nil {
		return nil, err
	}

	limit := req.PageSize
	if limit == 0 {
		limit = defaultPageSize
	}

	result, moreResults, err := s.feesService.ListFees(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("got error from fees service: %w", err)
	}

	var pageToken string
	if moreResults {
		pageToken = common.EncodeToken(offset + limit)
	}

	return &feespb.ListFeesResponse{
		Fees:              ToExternalFees(result),
		PreviousPageToken: common.GetPreviousPageToken(offset, limit),
		NextPageToken:     pageToken,
	}, err
}
