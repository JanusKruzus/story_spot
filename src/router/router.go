package router

import (
	"fmt"
	"net/http"
	"storyspot/src/auth"
	"storyspot/src/db"

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
	r.GET("/admin", h.admin)

	r.POST("/login", h.loginForm)
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

func (h *handler) admin(c *gin.Context) {

	var users []db.User
	db.GetTableRows(&users)

	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"users": users,
	})
}

/*	fmt.Println("Everything good to go 👍")
	fmt.Printf("name: %v\nemail: %v\npassword: %v\n", name, email, password)*/

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

	hash := auth.CreatePasswordHash(password)

	db.AddTable(db.User{
		Name:     name,
		Email:    email,
		Password: hash,
	})

}

func (h *handler) loginForm(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user db.User
	db.GetSingleRowsWhere(&user, "Email", email)

	if user.ID == 0 {
		fmt.Println("Invalid email")
		return
	}

	if auth.ComparePasswordHash(password, user.Password) != nil {
		fmt.Println("Invalid password")
		return
	}

	fmt.Println("User authenticated")
}
