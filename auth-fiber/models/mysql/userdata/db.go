package userdata

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// Mysql connection driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserhost     = "MYSQL_USERS_HOST"
	mysqlUserpassword = "MYSQL_USERS_PASSWORD"
	mysqlUserschema   = "MYSQL_USERS_SCHEMA"
	mysqlUserusername = "MYSQL_USERS_USERNAME"
)

// Client is the gorm db connection
var (
	Client *gorm.DB

	host     = os.Getenv(mysqlUserhost)
	password = os.Getenv(mysqlUserpassword)
	schema   = os.Getenv(mysqlUserschema)
	username = os.Getenv(mysqlUserusername)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		schema,
	)

	log.Println("Connecting to MySql DB")
	var err error
	Client, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.DB().Ping(); err != nil {
		panic(err)
	}

	log.Printf("MySql %s Database Connected", schema)

	Client.AutoMigrate(&User{})
	fmt.Println("Database Migrated")

}
