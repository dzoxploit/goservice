package migrations

import (
	"database/sql"
	"fmt"
	"go_service/database"
)

func User() {

	db := database.ConnectMigration()
	tableName := "users"
	// drop if exists
	qcheck := "DROP TABLE IF EXISTS %s"
	ExecQuery(db, fmt.Sprintf(qcheck, tableName))

	query := "CREATE TABLE %s (%s, %s, %s, %s, %s, %s)"

	value := []interface{}{
		tableName,
		"id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT",
		"username varchar(100) UNIQUE",
		"password varchar(100)",
		"email varchar(100)",
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP",
	}

	ExecQuery(db, fmt.Sprintf(query, value...))

	defer db.Close()
}

func ExecQuery(db *sql.DB, s string) {
	panic("unimplemented")
}