package repository

import (
	"context"
	"customer/domain"
	"customer/pkg"
	"errors"

	"gorm.io/gorm"
)

type postgresCustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &postgresCustomerRepository{db}
}

func (p *postgresCustomerRepository) Fetch(ctx context.Context, pg pkg.Pagination) (res []domain.Customer, nextpg pkg.Pagination, err error) {
	err = p.db.WithContext(ctx).Scopes(pkg.Paginate(domain.Customer{}, &pg, p.db)).Find(&res).Error
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = domain.ErrCustomerNoFound
		return
	}

	nextpg = pg

	return
}

func (p *postgresCustomerRepository) GetById(ctx context.Context, id int64) (res domain.Customer, err error) {
	err = p.db.WithContext(ctx).First(&res, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = domain.ErrCustomerNoFound
		return
	}

	if err != nil {
		return
	}

	return
}
