package domain

import (
	"context"
	"customer/pkg"
	"database/sql"
	"errors"
)

var (
	ErrCustomerNoFound = errors.New("customer not found")
)

type (
	CustomerUsecase interface {
		Fetch(context.Context, pkg.Pagination) ([]Customer, pkg.Pagination, error)
		GetById(context.Context, int64) (Customer, error)
	}

	CustomerRepository interface {
		Fetch(context.Context, pkg.Pagination) ([]Customer, pkg.Pagination, error)
		GetById(context.Context, int64) (Customer, error)
	}

	Customer struct {
		Customerid           int64 `gorm:"primaryKey"`
		Firstname            string
		Lastname             string
		City                 string
		State                sql.NullString
		Zip                  sql.NullInt64
		Country              string
		Region               int64
		Email                sql.NullString
		Phone                sql.NullString
		Creditcardtype       int64
		Creditcard           string
		Creditcardexpiration string
		Username             string
		Password             string
		Age                  sql.NullInt64
		Income               sql.NullInt64
		Gender               sql.NullString
	}
)
