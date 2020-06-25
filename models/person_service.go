package models

import (
	"Neo4jCURD/config"
	"fmt"
	"log"
)

var (
	PersonList map[string]*Person
)

type Persons struct {
	Counts      int       `json:"counts"`
	ListPersons []*Person `json:"persons"`
}

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ReqPerson struct {
	Node  string `json:"node"`
	Label string `json:"label"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

func getNodePerson(personId string) (*Person, error) {
	query := `
		MATCH (p:Person)
		WHERE p.id = toInt($personid)
		RETURN
			p.id as id,
			p.name as name`

	mapParams := map[string]interface{}{
		"personid": personId,
	}

	result, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return nil, err
	}

	var temp_Person *Person
	if result.Next() {
		id := fmt.Sprintf("%v", result.Record().GetByIndex(0))
		name := fmt.Sprintf("%v", result.Record().GetByIndex(1))

		temp_Person = &Person{
			id,
			name,
		}
	}

	return temp_Person, nil
}

func getAllNodePersons() map[string]*Person {
	PersonList := make(map[string]*Person)

	query := `
		MATCH (p:Person)
		RETURN
			p.id as id,
			p.name as name`

	mapParams := map[string]interface{}{}
	result, err := config.ResultQuery(query, mapParams)
	if err != nil {
		log.Println("", err)
		return nil
	}

	for result.Next() {
		id := fmt.Sprintf("%v", result.Record().GetByIndex(0))
		name := fmt.Sprintf("%v", result.Record().GetByIndex(1))

		temp_Person := &Person{
			id,
			name,
		}

		PersonList[id] = temp_Person
	}

	return PersonList
}

func editNodePerson(personId string, p *Person) (string, error) {
	query := `
			MATCH (p:Person)
			WHERE 
				p.id = toInt($personid)
			SET
				p.name = $name`

	mapPrams := map[string]interface{}{
		"personid": personId,
		"name":    p.Name,
	}

	_, err := config.ResultQuery(query, mapPrams)
	if err != nil {
		return "false", err
	}

	return "true", nil
}

func deleteNodePerson(personId string) (error)  {
	query := fmt.Sprintf("MATCH (p:Person) WHERE p.id = %s DETACH DELETE p",
		personId)
	mapParams := map[string]interface{}{}

	_, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return err
	}

	return nil
}

func addNodePerson(reqperson *ReqPerson) (error)  {
	query := fmt.Sprintf("CREATE (%s:%s {id: %s, name: \"%s\"})",
		reqperson.Node, reqperson.Label, reqperson.Id, reqperson.Name)

	mapParam := map[string]interface{}{}

	_, err := config.ResultQuery(query, mapParam)
	if err != nil {
		return err
	}

	return nil
}
