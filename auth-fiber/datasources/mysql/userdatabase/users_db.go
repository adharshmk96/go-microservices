package userdatabase

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

// User describes a User
type User struct {
	gorm.Model
	FirstName string `json:"first_name"	form:"first_name"		query:"first_name"`
	LastName  string `json:"last_name" 	form:"last_name"		query:"last_name"`
	Email     string `json:"email" 		form:"email"		 	query:"email"			gorm:"type:varchar(100);unique_index"`
}

// Client is the gorm db connection
var (
	Client *gorm.DB

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

	log.Println("Connecting to MySql DB")
	var err error
	Client, err = gorm.Open("mysql", dataSourceName)
	defer Client.Close()

	if err != nil {
		panic(err)
	}

	if err = Client.DB().Ping(); err != nil {
		panic(err)
	}

	log.Println("MySql gin_user_db Database Connected")

	Client.AutoMigrate(&User{})
	fmt.Println("Database Migrated")

}
