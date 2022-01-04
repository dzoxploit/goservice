package models

import (
	"database/sql"
	"fmt"
	"go_service/app/models/entities"

	_ "github.com/go-sql-driver/mysql"
)


type PostModel struct {
	Db *sql.DB
}


func (postModel PostModel) FindPost(slug string) entities.PostEntity {

	db := postModel.Db

	var post = entities.PostEntity{}

	rows,err := db.QueryRow(`
				SELECT id,
				title, 
				content, 
				image
				FROM posts WHERE slug=?
				`, slug).
					Scan(
						&post.ID,
						&post.Title,
						&post.Content,
						&post.Image,
						&post.Created_at,
					)

	if err != nil {
		fmt.Sprint()
	}

	defer rows.Close()
	

	
	return post
}
