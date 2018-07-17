package storage

import (
	"database/sql"
)

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS blogs(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR NOT NULL,
		url VARCHAR NOT NULL,
		published DATETIME
	);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
