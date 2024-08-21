package main

import (
	"fmt"
	"storyspot/src/db"
	"storyspot/src/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db.DbConnect()
	db.DbMigrate()

	var users []db.User
	db.GetTableRows(&users) //inserts into users selected rows from of the same data type

	for i, user := range users {
		fmt.Printf("%v. query result: %v\n", i, user.Name)
	}

	gin.SetMode(gin.ReleaseMode)

	router.SetupRoutes()

}
