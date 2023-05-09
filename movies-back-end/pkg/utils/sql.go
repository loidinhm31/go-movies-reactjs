package utils

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
