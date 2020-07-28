package dbManager

import (
    "fmt"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "default/models/logs"
)

func ConnectDatabase( dbName string ) *sql.DB  {
    dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s" , os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD") , os.Getenv("DB_HOST") , os.Getenv("DB_PORT") , dbName )
	db, err := sql.Open("mysql", dbConnection)
	logs.PutError( 1 , err )
	return db
}
