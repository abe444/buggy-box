package routes

import (
	controllers "buggybox/controllers"

	"github.com/gin-gonic/gin"
)

func ServeNRoute(r *gin.Engine) {

	// Serve CSS and favicon
	r.StaticFile("/favicon", "views/static/images/favicon.ico")
	r.StaticFile("/styles.css", "views/static/styles.css")

	// Serve HTML templates
	r.LoadHTMLGlob("views/templates/*")

	r.GET("/", controllers.Index)

	r.POST("login", controllers.Login)
	r.POST("upload", controllers.UploadFile)
}
