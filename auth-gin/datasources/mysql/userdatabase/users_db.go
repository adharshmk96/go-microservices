package userdatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Mysql connection driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserhost     = "MYSQL_USERS_HOST"
	mysqlUserpassword = "MYSQL_USERS_PASSWORD"
	mysqlUserschema   = "MYSQL_USERS_SCHEMA"
	mysqlUserusername = "MYSQL_USERS_USERNAME"
)

// UserDB is the mysql connection
var (
	Client *sql.DB

	host     = os.Getenv(mysqlUserhost)
	password = os.Getenv(mysqlUserpassword)
	schema   = os.Getenv(mysqlUserschema)
	username = os.Getenv(mysqlUserusername)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	log.Println("Connecting to Mysql DB...")
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	// mysql.SetLogger()

	log.Println("MySql gin_user_db Database Connected")
}
