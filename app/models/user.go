package models

import (
	"database/sql"
	"go_service/app/models/entities"

	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	Db *sql.DB
}


func (userModel UserModel) FindUser(username string) entities.UserEntity {

	db := userModel.Db

	var user = entities.UserEntity{}

	db.QueryRow(`
	SELECT id,
	name, 
	username, 
	email, 
	password 
	FROM users WHERE username=?
	`, username).
		Scan(
			&user.Id_user,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Status,
		)

	defer db.Close()

	return user
}