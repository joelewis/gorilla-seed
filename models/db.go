package models

import (
	"github.com/jinzhu/gorm"
	utils "github.com/joelewis/gorilla-seed/utils"
	"log"
)

var Db *gorm.DB

func checkErr(e error, msg string) {
	if e != nil {
		log.Printf("%s %s", msg, e)
	}
}

func InitDb(db_url, dialect string) *gorm.DB {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := gorm.Open(dialect, db_url)
	db.LogMode(true)
	checkErr(err, "sql.Open failed")
	Db = db
	return db
}

func SetDb(db *gorm.DB) {
	Db = db
}

func GetDb() *gorm.DB {
	return Db
}

func InitTables(db_url, db_dialect string) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := gorm.Open(db_dialect, db_url)
	utils.CheckErr(err, "sql.Open failed")

	db.AutoMigrate(&User{})
	db.Close()
}
