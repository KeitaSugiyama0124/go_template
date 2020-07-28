package contentsInfo

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
	"default/models/logs"
)

type Params struct {
    Id int
    ContentKey string
}

var tableName string = "contents_info"

func isRecordFound( db *sql.DB , option string ) bool {

	query := fmt.Sprintf( "select count(*) from %s %s" , tableName , option )
	rows, err := db.Query( query )
	logs.PutError( 1 , err )
	
	count := 0
	for rows.Next() {
		err := rows.Scan( 
			&count ,
		)
		logs.PutError( 1 , err )
	}

	if count > 0 {
		return true
	} else {
		logs.PutString( 2 , fmt.Sprintf( " table[%s] search record not found. serchOption ::: %s" , tableName , option ) )
		return false
	}
}

func genelateRowToParams( rows *sql.Rows) []Params {
	s := make( []Params , 0 )
	for rows.Next() {
		var params Params
		err := rows.Scan( 
			&params.Id , 
			&params.ContentKey , 
		)
		logs.PutError( 1 , err )
		s = append( s , params )
	}
	defer rows.Close()
	return s
}

func FindAll( db *sql.DB ) []Params {

	rows, err := db.Query( fmt.Sprintf( "select * from %s" , tableName ) )
    logs.PutError( 1 , err )
	return genelateRowToParams( rows )	
}
func Find( db *sql.DB , option string ) []Params {

	query := fmt.Sprintf( "select * from %s %s" , tableName , option )
	rows, err := db.Query( query )
    logs.PutError( 1 , err )

	return genelateRowToParams( rows )	
}

func Insert( db *sql.DB , params Params ) {

	columns := 
		"content_key"

	values := 
		"'" + params.ContentKey + "'"

	query := fmt.Sprintf( "insert into %s(%s) values(%s)" , tableName , columns , values )
	result, err := db.Query( query )
	defer result.Close()
	logs.PutError( 1 , err )
}
func Delete( db *sql.DB , option string ) {

	if !isRecordFound( db , option ) {
		return
	}

	query := fmt.Sprintf( "delete from %s %s" , tableName , option )
	result, err := db.Query( query )
	defer result.Close()
	logs.PutError( 1 , err )
}
func Update( db *sql.DB , option string , changedValue Params ) {

	if !isRecordFound( db , option ) {
		return
	}

	values :=
		fmt.Sprintf( "content_key='%s'" , changedValue.ContentKey )

	query := fmt.Sprintf( "update %s set %s %s" , tableName , values , option )
	logs.PutString( 1 , query )
	result, err := db.Query( query )
	defer result.Close()
	logs.PutError( 1 , err )
}
