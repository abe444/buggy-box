package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./uploads/"+file.Filename)

	c.String(http.StatusOK, "Uploaded %s successfully", file.Filename)
}
