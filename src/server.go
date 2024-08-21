package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"size:64"`
	Email    string `gorm:"size:64"`
	Password string `gorm:"size:255"`
}

var db *gorm.DB
var db_err error

func dbConnect() {
	db, db_err = gorm.Open(sqlite.Open("wsqlite.db"), &gorm.Config{})

	if db_err != nil {
		panic("Failed to connect to the sqlite database")
	}
}

func dbMigrate() {
	db.AutoMigrate(&User{})
}

func main() {

	dbConnect()
	dbMigrate()

	/*user1 := User{
		Name:     "Ambatukam",
		Email:    "basin@gmail.com",
		Password: "bust",
	}

	result := db.Create(&user1)
	fmt.Printf("result: %v\n", result.Error)*/

	var users []User
	qResult := db.Find(&users)

	if qResult.Error != nil {
		fmt.Println("Error fetching users:", qResult.Error)
		return
	}

	for i, user := range users {
		fmt.Printf("%v. query result: %v\n", i, user.Name)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static/")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "main website",
		})
	})

	fmt.Printf("Server running on localhost:3000...\n")
	router.Run(":3000")

}
