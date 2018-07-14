package mysql

import (
	"database/sql"
	"errors"

	"github.com/gedelumbung/go-movie/model"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

const selectCategory = `select id, name, created_at, updated_at from categories`

func (o *categoryRepository) All(page, limit int) ([]model.Category, int, error) {
	var (
		categories   []model.Category
		count, start int
	)
	categories = []model.Category{}

	err := o.db.QueryRowx(`select count(*) from categories where deleted_at is null`).Scan(&count)
	if err != nil {
		return categories, 0, err
	}

	start = (page - 1) * limit
	if start < 0 || limit < 1 {
		return categories, 0, errors.New("insufficient parameters")
	}

	err = o.db.Select(&categories, selectCategory+` limit ? offset ?`, limit, start)

	return categories, count, nil
}

func (o *categoryRepository) FindByID(id int) (model.Category, error) {
	var category model.Category
	err := o.db.QueryRowx(selectCategory+` where id = ?`, id).StructScan(&category)
	if err == sql.ErrNoRows {
		return category, repository.ErrNotFound
	}
	if err != nil {
		return category, err
	}
	return category, err
}
