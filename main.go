package main

import (
	"buggybox/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.ServeNRoute(r)

	r.Run(":3000")
}
