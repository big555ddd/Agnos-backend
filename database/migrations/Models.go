package migrations

import "app/app/model"

func Models() []any {
	return []any{
		(*model.Staff)(nil),
		(*model.Patient)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
