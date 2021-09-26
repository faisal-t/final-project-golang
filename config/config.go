package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "sql6439664"
	password string = "Pae356TDXy"
	host     string = "sql6.freesqldatabase.com"
	port     int    = 3306
	database string = "sql6439664"
)

var (
	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
)

// const (
// 	username string = "root"
// 	password string = ""
// 	database string = "final-project-golang"
// )

// var (
// 	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
// )

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
