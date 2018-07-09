package helper

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NullInt64ToInt(i sql.NullInt64) int {
	if i.Valid {
		return int(i.Int64)
	}
	return 0
}

func NullStringToString(i sql.NullString) string {
	if i.Valid {
		return i.String
	}
	return ""
}

func NullTimeToString(i mysql.NullTime, format string) string {
	if i.Valid {
		return i.Time.Format(format)
	}
	return ""
}

func NullFloatToString(i sql.NullFloat64) string {
	if i.Valid {
		return strconv.FormatFloat(i.Float64, 'g', -1, 64)
	}
	return ""
}

func StringToNullString(i string) sql.NullString {
	return sql.NullString{
		String: i,
		Valid:  true,
	}
}

func TimeToNullTime(t time.Time) mysql.NullTime {
	return mysql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func IntToNullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func BoolToInt(i bool) int {
	if i {
		return 1
	}
	return 0
}

func ArrayIntToStringSeparatedByComma(arr []int) string {
	var val []string
	for _, i := range arr {
		val = append(val, strconv.Itoa(i))
	}
	return strings.Join(val, ", ")
}
