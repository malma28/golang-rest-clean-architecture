package repository

import (
	"context"

	"github.com/malma28/golang-rest-clean-architecture/entity"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	FindAll(ctx context.Context) ([]entity.Customer, error)
	FindById(ctx context.Context, id int64) (entity.Customer, error)
	UpdateById(ctx context.Context, id int64, customer entity.Customer) (entity.Customer, error)
	DeleteById(ctx context.Context, id int64) error
	WithTx(ctx context.Context, callback func(ctx context.Context) error) error
}
