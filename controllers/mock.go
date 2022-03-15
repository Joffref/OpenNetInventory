package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mock": "ok"})
}
