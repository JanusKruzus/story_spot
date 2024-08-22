package router

import (
	"fmt"
	"net/http"
	"storyspot/src/auth"

	"github.com/gin-gonic/gin"
)

type handler struct{}

var r *gin.Engine

func SetupRoutes() {

	h := &handler{}
	r = gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static/")

	r.GET("/home", h.home)
	r.GET("/login", h.login)
	r.GET("/signup", h.signup)

	r.POST("/signup", h.signupForm)

	fmt.Printf("Server running on localhost:3000...\n")
	r.Run(":3000")
}

func (h *handler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "idk",
	})
}

func (h *handler) login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "idk",
	})
}

func (h *handler) signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{
		"title": "idk",
	})
}

func (h *handler) signupForm(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if auth.CheckEmail(email) != nil {
		fmt.Println("Error in email validation")
		return
	}

	if auth.CheckPassword(password) != nil {
		fmt.Println("Error in password validation")
		return
	}

	fmt.Println("Everything good to go 👍")
	fmt.Printf("name: %v\nemail: %v\npassword: %v\n", name, email, password)
}
