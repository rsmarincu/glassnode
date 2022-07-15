package grpc

import (
	feespb "github.com/rsmarincu/glassnode/api"
	"github.com/rsmarincu/glassnode/pkg/fees"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToExternalFee(internalFee *fees.Fee) *feespb.Fee {
	return &feespb.Fee{
		T: timestamppb.New(internalFee.Timestamp),
		V: internalFee.Value,
	}
}

func ToExternalFees(internalFees []*fees.Fee) []*feespb.Fee {
	protoFees := make([]*feespb.Fee, len(internalFees))
	for i, fee := range internalFees {
		protoFees[i] = ToExternalFee(fee)
	}
	return protoFees
}
