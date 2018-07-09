package helper

import (
	"github.com/go-sql-driver/mysql"
)

func NullTimeToString(i mysql.NullTime, format string) string {
	if i.Valid {
		return i.Time.Format(format)
	}
	return ""
}
