package datasources

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

const (
	db_user   = "root"
	db_pass   = "9a147gml"
	db_host   = "127.0.0.1:3306"
	db_schema = "Authors"
)

func init() {
	var err error
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s", db_user, db_pass, db_host, db_schema)
	Client, err = sql.Open("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	if err = Client.Ping(); err != nil {
		fmt.Println("Es em qaqmech anum")
		log.Fatal(err)
	}

	fmt.Println("successfully connected to db")
}
