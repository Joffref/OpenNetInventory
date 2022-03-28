package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joffref/openNetInventory/db"
	"github.com/joffref/openNetInventory/models"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetAllNetwork(c *gin.Context, driver neo4j.Driver) {
	networks := []models.Network{}
	res, err := db.ReadInDb(driver, "MATCH (n:Network) RETURN n")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	mapstructure.Decode(res, &networks)
	c.JSON(http.StatusOK, gin.H{"networks": networks})
}

func GetNetworkById(c * gin.Context, drive neo4j;driver) {
	
}