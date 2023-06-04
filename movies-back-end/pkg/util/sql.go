package util

import "database/sql"

func StringToSQLNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func FloatToSQLNullFloat(s float64) sql.NullFloat64 {
	if s != 0 {
		return sql.NullFloat64{
			Float64: s,
			Valid:   true,
		}
	}
	return sql.NullFloat64{}
}

func IntToSQLNullInt(s int64) sql.NullInt64 {
	if s != 0 {
		return sql.NullInt64{
			Int64: s,
			Valid: true,
		}
	}
	return sql.NullInt64{}
}
