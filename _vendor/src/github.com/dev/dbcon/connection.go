package dbcon

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "eddifisher"
	password = ""
	dbname   = "postblog"
	sslmode  = "disable"
)

func Connection() *sql.DB {
	var dbinfo string
	if password == "" {
		dbinfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", host, port, user, dbname, sslmode)
	} else {
		dbinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	}

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	logrus.Info("Successfully connected to postgres!")
	return db
}

func checkErr(err error) {
	if err != nil {
		logrus.Panicf("%v", err)
	}
}
