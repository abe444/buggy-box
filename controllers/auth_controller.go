package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("username")

	query := "SELECT * FROM users WHERE username='" + username + "' AND password='" + password + "'"

	c.String(http.StatusOK, "Executing SQL Query: %s", query)
}
