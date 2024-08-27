package router

import (
	"fmt"
	"net/http"
	"os"
	"storyspot/src/auth"
	"storyspot/src/db"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type Handler struct{}

var store *sessions.CookieStore

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	store = sessions.NewCookieStore([]byte(os.Getenv("AUTH_KEY")))
}

func Home(c *gin.Context) {
	session, _ := store.Get(c.Request, "login")

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"name": session.Values["name"],
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "idk",
	})
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{
		"title": "idk",
	})
}

func Admin(c *gin.Context) {

	var users []db.User
	db.GetTableRows(&users)

	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"users": users,
	})
}

func SignupForm(c *gin.Context) {
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

func LoginForm(c *gin.Context) {
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

	session, errSession := store.Get(c.Request, "login")

	if errSession != nil {
		c.String(http.StatusInternalServerError, "Unable to store session")
		return
	}

	session.Values["name"] = user.Name

	errSave := session.Save(c.Request, c.Writer)
	if errSave != nil {
		c.String(http.StatusInternalServerError, "Unable to save session")
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}
