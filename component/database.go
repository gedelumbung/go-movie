package component

import (
	"fmt"

	"github.com/gedelumbung/go-movie/config"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/gedelumbung/go-movie/repository/mysql"
	"github.com/gedelumbung/go-movie/repository/sqlite3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func GetDatabaseConnection(config *conf.Configuration) (repository.Repository, error) {
	if config.DB.Driver == "mysql" {
		return mysql.Connect(config.DB.Mysql.URL)
	}
	if config.DB.Driver == "sqlite3" {
		return sqlite3.Connect(config.DB.Sqlite3.URL)
	}
	return nil, fmt.Errorf("unknown store type: %s", config.DB.Driver)
}
