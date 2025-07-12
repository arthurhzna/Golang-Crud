package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api_golang/helper"
	"rest_api_golang/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) VALUES ($1) RETURNING id "
	var id int
	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&id) 
	helper.ErrorConditionCheck(err)

	category.Id = id
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = $1 WHERE id = $2"
	_ , err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.ErrorConditionCheck(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.ErrorConditionCheck(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.ErrorConditionCheck(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorConditionCheck(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.ErrorConditionCheck(err)

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorConditionCheck(err)
		categories = append(categories, category)
	}
	return categories
}
