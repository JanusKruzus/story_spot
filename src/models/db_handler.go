package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbConnect(db *gorm.DB, db_err error) {
	db, db_err = gorm.Open(sqlite.Open("wsqlite.db"), &gorm.Config{})

	if db_err != nil {
		panic("Failed to connect to the sqlite database")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

//func addOneToTable(){}
