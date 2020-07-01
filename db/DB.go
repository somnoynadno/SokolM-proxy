package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

var username string
var password string
var dbName   string
var dbHost   string
var dbPort   string
var dbURI    string

func init() {
	username = os.Getenv("db_user")
	password = os.Getenv("db_pass")
	dbName = os.Getenv("db_name")
	dbHost = os.Getenv("db_host")
	dbPort = os.Getenv("db_port")

	dbURI = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, dbHost + ":" + dbPort, dbName)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	} else {
		db = conn
	}

	createTables()
}

func createTables() {
	conn, err := sql.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec("USE apik3;")

	for _, d := range Devices {
		query := "CREATE TABLE IF NOT EXISTS `" + d.NumberInDB + "` ( " +
			"`time` int(10) unsigned NOT NULL DEFAULT '0', " +
			"EVS  float, " +
			"UVI  float, " +
			"L    float, " +
			"LI   float, " +
			"RSSI float, " +
			"RN   float, " +
			"T    float, " +
			"WD   float, " +
			"HM   float, " +
			"WV   float, " +
			"WM   float, " +
			"UV   float, " +
			"Upow float, " +
			"PR1  float, " +
			"PR   float, " +
			"KS   float, " +
			"TR   float, " +
			"TD   float, " +
			"PRIMARY KEY (`time`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=latin1;"
		_, err := conn.Exec(query)

		if err != nil {
			panic(err)
		}
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
