package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joffref/netrman/models"
)

func GetPortById(c *gin.Context) {
	id := c.Param("id")
	MyPort := new(models.Port)
	MyPort.Name = id
	c.JSON(http.StatusOK, MyPort)
}
