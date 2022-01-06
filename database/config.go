package database

import (
	"os"
)

type dbconfig struct {
	Driver string
	User string
	Pass string
	Name string
	Host string
}

func config() dbconfig {
	_db := dbconfig{
		Driver: os.Getenv("DB_CONNECTION"),
		User: os.Getenv("DB_USERNAME"),
		Pass: os.Getenv("DB_PASSWORD"),
		Name: os.Getenv("DB_DATABASE"),
		Host: os.Getenv("DB_HOST"),
	}

	return _db
}

