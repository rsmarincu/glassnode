package usecases

import (
	"context"
	"math"
	"time"

	fees "github.com/rsmarincu/glassnode/pkg/fees"
	"github.com/rsmarincu/glassnode/pkg/fees/repository"
)

//mockgen -destination pkg/usecases/mocks/ethRepositoryMock.go github.com/rsmarincu/glassnode/pkg/usecases ETHRepository
type ETHRepository interface {
	QueryEOATransactions(ctx context.Context) ([]*repository.Transaction, error)
}

type FeesService struct {
	ETHRepository ETHRepository
}

func NewFeesService(repo ETHRepository) *FeesService {
	return &FeesService{
		ETHRepository: repo,
	}
}

func (f *FeesService) ListFees(ctx context.Context, offset uint32, limit uint32) ([]*fees.Fee, bool, error) {
	transactions, err := f.ETHRepository.QueryEOATransactions(ctx)
	if err != nil {
		return nil, false, err
	}

	if len(transactions) == 0 {
		return nil, false, nil
	}

	finalFees := []*fees.Fee{}
	currentHour := transactions[0].BlockTime.Truncate(60 * time.Minute)
	nextHour := currentHour.Add(60 * time.Minute)

	var currentHourValue float64
	for _, t := range transactions {
		if t.BlockTime.Before(nextHour) {
			currentHourValue += t.GasPrice * t.GasUsed
		} else {
			finalFees = append(finalFees, &fees.Fee{
				Timestamp: currentHour,
				Value:     currentHourValue * math.Pow(10, -18),
			})
			currentHourValue = 0
			nextHour = nextHour.Add(60 * time.Minute)
		}
	}

	if offset > 0 {
		finalFees = finalFees[offset:]
	}

	var moreResults bool
	if moreResults = len(finalFees) > int(limit); moreResults {
		finalFees = finalFees[:limit]
	}

	return finalFees, moreResults, nil
}
