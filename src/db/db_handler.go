package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbError error

func DbConnect() {
	Db, DbError = gorm.Open(sqlite.Open("wsqlite.db"), &gorm.Config{})

	if DbError != nil {
		panic("Failed to connect to the sqlite database")
	}

	fmt.Println("-> Connected to database")
}

func DbMigrate() {
	Db.AutoMigrate(&User{})
}

func AddTable[T User](table T) {
	result := Db.Create(&table)
	if result.Error != nil {
		fmt.Println("Error inserting: ", result.Error)
	}
}

func GetTableRows[T User](table *[]T) { //inserts into table selected rows of the same data type
	result := Db.Find(table)
	if result.Error != nil {
		fmt.Println("Error fetching: ", result.Error)
	}
}
