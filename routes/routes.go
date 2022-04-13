package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joffref/openNetInventory/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/status", controllers.GetStatus)
		port := v1.Group("/port")
		{
			port.GET("/", controllers.Mock)
			port.POST("/", controllers.Mock)

			subport := v1.Group("/subport")
			{
				subport.GET("/", controllers.Mock)
				subport.POST("/", controllers.Mock)
			}
		}
		network := v1.Group("/network")
		{
			network.GET("/", controllers.GetAllNetworks)
			network.POST("/", controllers.CreateNetwork)
			network.PUT("/", controllers.UpdateNetwork)
			network.DELETE("/", controllers.DeleteNetwork)
			ip := network.Group("/ip")
			{
				ip.DELETE("/:ip", controllers.DeleteNetworkByIP)
				ip.GET("/:ip", controllers.GetNetworkByIP)
			}
			id := network.Group("/id")
			{
				id.DELETE("/:id", controllers.DeleteNetworkByID)
				id.GET("/:id", controllers.GetNetworkByID)
			}
		}
		link := v1.Group("/link")
		{
			link.GET("/", controllers.Mock)
			link.POST("/", controllers.Mock)
		}
		equipement := v1.Group("/equipement")
		{
			equipement.GET("/", controllers.Mock)
			equipement.POST("/", controllers.Mock)
		}
		cluster := v1.Group("/cluster")
		{
			cluster.GET("/", controllers.Mock)
			cluster.POST("/", controllers.Mock)
		}
	}

	return r
}
