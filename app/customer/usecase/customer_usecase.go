package usecase

import (
	"context"
	"customer/domain"
	"customer/pkg"
	"time"
)

type CustomerUsecase struct {
	ctxTimeout   time.Duration
	customerRepo domain.CustomerRepository
}

func NewCustomerUsecase(to time.Duration, c domain.CustomerRepository) domain.CustomerUsecase {
	return &CustomerUsecase{
		ctxTimeout:   to,
		customerRepo: c,
	}
}

func (uc *CustomerUsecase) Fetch(ctx context.Context, pg pkg.Pagination) (res []domain.Customer, nextpg pkg.Pagination, err error) {
	c, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	res, nextpg, err = uc.customerRepo.Fetch(c, pg)
	return
}

func (uc *CustomerUsecase) GetById(ctx context.Context, id int64) (res domain.Customer, err error) {
	c, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	res, err = uc.customerRepo.GetById(c, id)
	return
}
