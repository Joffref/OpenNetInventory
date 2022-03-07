package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joffref/netrman/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/status", controllers.GetStatus)
		port := v1.Group("/port")
		{
			port.GET("/:id", controllers.GetPortById)
		}
	}

	return r
}
