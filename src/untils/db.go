package untils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	e  error
)

func init() {

	Db, e = sql.Open("mysql", "qmpo:qmpo_1234@tcp(10.133.206.139:3306)/qmpo_qmf")
	if e != nil {
		panic(e.Error())
	}
	Db.Ping()
}

