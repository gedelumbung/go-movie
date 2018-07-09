package repository

import (
	"errors"

	"github.com/gedelumbung/go-movie/model"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Repository interface {
	Categories() CategoryRepository
}

type CategoryRepository interface {
	FindByID(id int) (model.Category, error)
}
