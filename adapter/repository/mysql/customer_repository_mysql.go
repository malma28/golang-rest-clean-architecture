package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/malma28/golang-rest-clean-architecture/adapter/database"
	"github.com/malma28/golang-rest-clean-architecture/adapter/repository"
	"github.com/malma28/golang-rest-clean-architecture/entity"
	"github.com/malma28/golang-rest-clean-architecture/entity/exception"
)

type customerRepositoryMysql struct {
	db database.SQL
}

func NewCustomerRepository(db database.SQL) repository.CustomerRepository {
	customerRepository := new(customerRepositoryMysql)
	customerRepository.db = db

	return customerRepository
}

func (repository *customerRepositoryMysql) withTx(ctx context.Context, callback func(tx database.Tx) error) (err error) {
	tx, isTx := ctx.Value("tx").(database.Tx)
	if tx == nil || !isTx {
		tx, err = repository.db.Tx(ctx)
		if err != nil {
			return
		}
		defer func() {
			errTxCommit := tx.Commit()
			if err == nil {
				err = errTxCommit
			}
		}()
	}

	return callback(tx)
}

func (repository *customerRepositoryMysql) Save(ctx context.Context, customer entity.Customer) (res entity.Customer, err error) {
	res = customer

	err = repository.withTx(ctx, func(tx database.Tx) error {
		execRes, err := tx.Exec(
			ctx,
			`INSERT INTO customer (name) VALUES (?);`,
			customer.Name,
		)
		if err != nil {
			return err
		}

		res.Id, err = execRes.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (repository *customerRepositoryMysql) FindAll(ctx context.Context) ([]entity.Customer, error) {
	customers := []entity.Customer{}

	err := repository.withTx(ctx, func(tx database.Tx) error {
		rows, err := tx.Query(ctx, "SELECT * FROM customer;")
		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			var customer entity.Customer
			if err := rows.Scan(&customer.Id, &customer.Name); err != nil {
				return err
			}
			customers = append(customers, customer)
		}

		return rows.Err()
	})

	return customers, err
}

func (repository *customerRepositoryMysql) FindById(ctx context.Context, id int64) (entity.Customer, error) {
	var customer entity.Customer

	err := repository.withTx(ctx, func(tx database.Tx) error {
		row, err := tx.QueryRow(ctx, "SELECT * FROM customer WHERE id = ?;", id)
		if err != nil {
			return err
		}

		if err := row.Scan(&customer.Id, &customer.Name); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return exception.ErrCustomerNotFound
			}
			return err
		}
		return nil
	})

	return customer, err
}

func (repository *customerRepositoryMysql) UpdateById(ctx context.Context, id int64, customer entity.Customer) (entity.Customer, error) {
	res := customer

	err := repository.withTx(ctx, func(tx database.Tx) error {
		execRes, err := tx.Exec(
			ctx,
			`UPDATE customer SET name = ? WHERE id = ?;`,
			customer.Name,
			id,
		)
		if err != nil {
			return err
		}

		rowsAffected, err := execRes.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected < 1 {
			return exception.ErrCustomerNotFound
		}

		res.Id = id

		return nil
	})

	return res, err
}

func (repository *customerRepositoryMysql) DeleteById(ctx context.Context, id int64) error {
	return repository.withTx(ctx, func(tx database.Tx) error {
		execRes, err := tx.Exec(
			ctx,
			`DELETE FROM customer WHERE id = ?;`,
			id,
		)
		if err != nil {
			return err
		}

		rowsAffected, err := execRes.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected < 1 {
			return exception.ErrCustomerNotFound
		}

		return nil
	})
}

func (repository *customerRepositoryMysql) WithTx(ctx context.Context, callback func(ctx context.Context) error) error {
	tx, err := repository.db.Tx(ctx)
	if err != nil {
		return err
	}

	if err = callback(context.WithValue(ctx, "tx", tx)); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
