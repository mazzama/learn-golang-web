package repository

import (
	"context"
	"database/sql"
	"mazzama/learn-golang-web/application/helper"
	"mazzama/learn-golang-web/application/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) domain.Product
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}

type ProductRepositoryImpl struct {
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Id)
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "select id, name from product where id = ?"
	row, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer row.Close()

	product := domain.Product{}
	if row.Next() {
		err := row.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, err
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select id, name from product"
	row, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer row.Close()

	var products []domain.Product
	if row.Next() {
		product := domain.Product{}
		err := row.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}
