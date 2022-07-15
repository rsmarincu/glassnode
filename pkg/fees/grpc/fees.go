package grpc

import (
	"context"
	"fmt"
	feespb "github.com/rsmarincu/glassnode/api"
	"github.com/rsmarincu/glassnode/pkg/fees"
)

type FeesService interface {
	ListFees(ctx context.Context) ([]*fees.Fee, error)
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
	fmt.Println("CALLING FEES SERVICE FROM  GRPC")
	_, err := s.feesService.ListFees(ctx)
	return nil, err
}
