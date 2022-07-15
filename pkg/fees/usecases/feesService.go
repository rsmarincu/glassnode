package usecases

import (
	"context"
	"fmt"

	fees "github.com/rsmarincu/glassnode/pkg/fees"
	"github.com/rsmarincu/glassnode/pkg/fees/repository"
)

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

func (f *FeesService) ListFees(ctx context.Context) ([]*fees.Fee, error) {
	transactions, err := f.ETHRepository.QueryEOATransactions(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("TRANSACTIONS: ", transactions[0])
	return nil, nil
}
