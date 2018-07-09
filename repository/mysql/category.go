package mysql

import (
	"database/sql"
	"fmt"

	"github.com/gedelumbung/go-movie/model"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

const selectCategory = `select id, name, created_at, updated_at from categories`

func (o *categoryRepository) FindByID(id int) (model.Category, error) {
	var category model.Category
	fmt.Println(category)
	err := o.db.QueryRowx(selectCategory+` where id = ?`, id).StructScan(&category)
	if err == sql.ErrNoRows {
		return category, repository.ErrNotFound
	}
	if err != nil {
		return category, err
	}
	return category, err
}
