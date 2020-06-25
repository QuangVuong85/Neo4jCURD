package config

import (
	"github.com/astaxie/beego"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

// Connect to Neo4j use neo4j-go-driver
func Connect2Neo4j() (neo4j.Driver, neo4j.Session, error) {
	configForNeo4j := func(config *neo4j.Config) {
		config.Encrypted = true
	}

	driver, err := neo4j.NewDriver(
		beego.AppConfig.String("URI"),
		neo4j.BasicAuth(
			beego.AppConfig.String("USER"),
			beego.AppConfig.String("PASSWORD"), ""),
		configForNeo4j)
	if err != nil {
		return nil, nil, err
	}

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, nil, err
	}

	return driver, session, nil
}

// result query node movies
func ResultQuery(query string, mapParams map[string]interface{}) (neo4j.Result, error) {
	driver, session, err := Connect2Neo4j()
	if err != nil {
		log.Println("Error connecting to Database: ", err)
	}
	defer driver.Close()
	defer session.Close()

	result, err := session.Run(query, mapParams)
	if err != nil {
		log.Println("", err)
		return nil, err
	}

	return result, err
}