package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func ConnectMigration() (db *sql.DB) {
	db, err := sql.Open(config().Driver, config().User+":"+config().Pass+"@tcp("+config().Host + ")" + "/"+config().Name)
	

	if err != nil {
		panic(err.Error())
	}

	return
}

func ConnectApi() *xorm.Engine {
	mysqlInfo := fmt.Sprintf("%s:%st@tcp(%s)/%s?charset=utf8", config().User, config().Pass, config().Host, config().Name)
	//format
	engine, err := xorm.NewEngine("", mysqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL() //It's necessary for Rookies

	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")
	return engine
}