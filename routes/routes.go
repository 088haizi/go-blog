package routes

import (
	"github.com/088haizi/go-blog/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	{
		v1.GET("/", controllers.Index)
		v1.POST("/posts", controllers.Create)
	}

	return r
}