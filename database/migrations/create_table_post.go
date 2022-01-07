package migrations

import (
	"fmt"
	"go_service/database"
)

func Post() {
	
	db := database.ConnectMigration()
	
	tableName := "posts"

	//drop if exists table

	qcheck := "DROP TABLE IF EXISTS %s"
	ExecQuery(db, fmt.Sprint(qcheck, tableName))

	query := "CREATE TABLE %s (%s, %s, %s, %s, %s, %s)"

	value := []interface{}{
		tableName, 
		"id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT",
		"title varchar(100) UNIQUE",
		"content text",
		"image varchar(100) NULL",
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
	}

	ExecQuery(db, fmt.Sprintf(query, value...))

	defer db.Close()
}

