package model

import (
	"encoding/json"
	"time"

	"github.com/gedelumbung/go-movie/helper"
	"github.com/go-sql-driver/mysql"
)

type Category struct {
	ID        int            `db:"id"`
	Name      string         `db:"name"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

func (o Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}{
		ID:        o.ID,
		Name:      o.Name,
		CreatedAt: helper.NullTimeToString(o.CreatedAt, time.RFC3339),
		UpdatedAt: helper.NullTimeToString(o.UpdatedAt, time.RFC3339),
		DeletedAt: helper.NullTimeToString(o.DeletedAt, time.RFC3339),
	})
}
