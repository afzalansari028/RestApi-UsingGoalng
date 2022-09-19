package helper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	DB_USER     = "postgres"
// 	DB_PASSWORD = "king"
// 	DB_NAME     = "library"
// )

// DB set up
func SetupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")

	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	return db
}
