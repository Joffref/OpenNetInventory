package db

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Connect() *neo4j.Driver {
	if os.Getenv("NEO4J_URL") == "" {
		os.Setenv("NEO4J_URL", "bolt://localhost:7687")
	}
	driver, err := neo4j.NewDriver(os.Getenv("NEO4J_URI"), neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), ""))
	if err != nil {
		panic(err)
	}
	return &driver
}

func Disconnect(driver neo4j.Driver) error {
	return driver.Close()
}

func InsertInDb(driver neo4j.Driver, req string, object map[string]interface{}) (interface{}, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(req, object)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	return result, err
}

func ReadInDb(driver neo4j.Driver, req string) (interface{}, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(req, nil)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	return result, err
}
