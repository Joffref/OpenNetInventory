package db

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Connect() neo4j.Driver {
	if os.Getenv("NEO4J_URL") == "" {
		os.Setenv("NEO4J_URL", "bolt://localhost:7687")
	}
	driver, err := neo4j.NewDriver(os.Getenv("NEO4J_URI"), neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), ""))
	if err != nil {
		panic(err)
	}
	return driver
}

func Disconnect(driver neo4j.Driver) error {
	return driver.Close()
}
