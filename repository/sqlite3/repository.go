package sqlite3

import (
	"github.com/gedelumbung/go-movie/repository"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db                 *sqlx.DB
	categoryRepository *categoryRepository
}

func (s *Repository) Categories() repository.CategoryRepository {
	return s.categoryRepository
}

var _ repository.Repository = (*Repository)(nil)

func Connect(url string) (*Repository, error) {
	db, err := sqlx.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s := &Repository{
		db:                 db,
		categoryRepository: &categoryRepository{db: db},
	}

	return s, nil
}
