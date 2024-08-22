package main

import (
	"storyspot/src/db"
	"storyspot/src/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db.DbConnect()
	db.DbMigrate()

	gin.SetMode(gin.ReleaseMode)

	router.SetupRoutes()

}
