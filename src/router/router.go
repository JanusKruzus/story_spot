package router

import (
	"fmt"
	"net/http"

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
