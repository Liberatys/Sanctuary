package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// TODO: add methods for different types of databases in order to have wider coverage.
func CreateDatabaseConnection(databaseInfo *DatabaseInformation) {
	db, err := sql.Open(databaseInfo.DatabaseTyp, fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", databaseInfo.Username, databaseInfo.Password, databaseInfo.IP, databaseInfo.Port, databaseInfo.DatabaseName))
	if err != nil {
		panic(err.Error())
	}
	database = db
}

func GetDatabaseConnection() *sql.DB {
	return database
}
