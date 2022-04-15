package controllers

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joffref/openNetInventory/db"
	"github.com/joffref/openNetInventory/models"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func CreateNetwork(c *gin.Context) {
	var network models.Network
	err := c.ShouldBindJSON(&network)
	network.ID = uuid.New().String()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(network.NetAddr)
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("CREATE (n:Network {id: $id, network_address: $network_address, mask: $mask, protocol: $protocol, gateway: $gateway,  tag_id: $tag_id, created_at: datetime(), updated_at: datetime()}) RETURN n",
			map[string]interface{}{
				"id":              network.ID,
				"network_address": network.NetAddr.String(),
				"mask":            network.Mask.String(),
				"protocol":        network.Protocol,
				"gateway":         network.Gateway.String(),
				"tag_id":          network.Tag,
			})
		if err != nil {
			return nil, err
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"network": records})
}

func GetAllNetworks(c *gin.Context) {
	var networks []models.Network
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) RETURN n", nil)
		if err != nil {
			return nil, err
		}
		for result.Next() {
			record := result.Record()
			network := models.Network{}
			network.ID = record.Values[0].(neo4j.Node).Props["id"].(string)
			network.NetAddr = net.ParseIP(record.Values[0].(neo4j.Node).Props["network_address"].(string))
			network.Mask = net.ParseIP(record.Values[0].(neo4j.Node).Props["mask"].(string))
			network.Protocol = record.Values[0].(neo4j.Node).Props["protocol"].(string)
			network.Gateway = net.ParseIP(record.Values[0].(neo4j.Node).Props["gateway"].(string))
			network.Tag = int(record.Values[0].(neo4j.Node).Props["tag_id"].(int64))
			network.CreatedAt = time.Time(record.Values[0].(neo4j.Node).Props["created_at"].(time.Time)).Format(time.RFC3339)
			network.UpdatedAt = time.Time(record.Values[0].(neo4j.Node).Props["updated_at"].(time.Time)).Format(time.RFC3339)
			log.Println(network)
			log.Println(network.CreatedAt)
			networks = append(networks, network)
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"networks": networks})
}

func GetNetworkByID(c *gin.Context) {
	var network models.Network
	id := c.Param("id")
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.id = $id RETURN n", map[string]interface{}{
			"id": id,
		})
		if err != nil {
			return nil, err
		}
		for result.Next() {
			record := result.Record()
			network.ID = record.Values[0].(neo4j.Node).Props["id"].(string)
			network.NetAddr = net.ParseIP(record.Values[0].(neo4j.Node).Props["network_address"].(string))
			network.Mask = net.ParseIP(record.Values[0].(neo4j.Node).Props["mask"].(string))
			network.Protocol = record.Values[0].(neo4j.Node).Props["protocol"].(string)
			network.Gateway = net.ParseIP(record.Values[0].(neo4j.Node).Props["gateway"].(string))
			network.Tag = int(record.Values[0].(neo4j.Node).Props["tag_id"].(int64))
			network.CreatedAt = time.Time(record.Values[0].(neo4j.Node).Props["created_at"].(time.Time)).Format(time.RFC3339)
			network.UpdatedAt = time.Time(record.Values[0].(neo4j.Node).Props["updated_at"].(time.Time)).Format(time.RFC3339)
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"network": network})
}

func GetNetworkByIP(c *gin.Context) {
	var network models.Network
	ip := c.Param("ip")
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.network_address = $ip RETURN n", map[string]interface{}{
			"ip": ip,
		})
		if err != nil {
			return nil, err
		}
		for result.Next() {
			record := result.Record()
			network.ID = record.Values[0].(neo4j.Node).Props["id"].(string)
			network.NetAddr = net.ParseIP(record.Values[0].(neo4j.Node).Props["network_address"].(string))
			network.Mask = net.ParseIP(record.Values[0].(neo4j.Node).Props["mask"].(string))
			network.Protocol = record.Values[0].(neo4j.Node).Props["protocol"].(string)
			network.Gateway = net.ParseIP(record.Values[0].(neo4j.Node).Props["gateway"].(string))
			network.Tag = int(record.Values[0].(neo4j.Node).Props["tag_id"].(int64))
			network.CreatedAt = time.Time(record.Values[0].(neo4j.Node).Props["created_at"].(time.Time)).Format(time.RFC3339)
			network.UpdatedAt = time.Time(record.Values[0].(neo4j.Node).Props["updated_at"].(time.Time)).Format(time.RFC3339)
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"network": network})
}

func UpdateNetwork(c *gin.Context) {
	var network models.Network
	err := c.ShouldBindJSON(&network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if network.ID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referencing by ID is not allowed"})
		return
	}
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.id = $id SET n.network_address = $network_address, n.mask = $mask, n.protocol = $protocol, n.gateway = $gateway, n.tag_id = $tag_id, n.updated_at = time() RETURN n",
			map[string]interface{}{
				"id":              network.ID,
				"network_address": network.NetAddr.String(),
				"mask":            network.Mask.String(),
				"protocol":        network.Protocol,
				"gateway":         network.Gateway.String(),
				"tag_id":          network.Tag,
			})
		if err != nil {
			return nil, err
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": records})
}

func DeleteNetwork(c *gin.Context) {
	var network models.Network
	err := c.ShouldBindJSON(&network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.id = $id DELETE n",
			map[string]interface{}{
				"id": network.ID,
			})
		if err != nil {
			return nil, err
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": records})
}

func DeleteNetworkByID(c *gin.Context) {
	id := c.Param("id")
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.id = $id DELETE n",
			map[string]interface{}{
				"id": id,
			})
		if err != nil {
			return nil, err
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": records})
}

func DeleteNetworkByIP(c *gin.Context) {
	ip := c.Param("ip")
	driver := db.Connect()
	defer db.Disconnect(driver)
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (n:Network) WHERE n.network_address = $ip DELETE n",
			map[string]interface{}{
				"ip": ip,
			})
		if err != nil {
			return nil, err
		}
		return result.Next(), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": records})
}
