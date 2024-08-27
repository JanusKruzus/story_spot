package main

import (
	"fmt"
	"storyspot/src/db"
	handler "storyspot/src/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	db.DbConnect()
	db.DbMigrate()
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static/")

	handler.Init()

	r.GET("/home", handler.Home)
	r.GET("/login", handler.Login)
	r.GET("/signup", handler.Signup)
	r.GET("/admin", handler.Admin)

	r.POST("/login", handler.LoginForm)
	r.POST("/signup", handler.SignupForm)

	fmt.Printf("Server running on localhost:3000...\n")
	r.Run(":3000")
}
