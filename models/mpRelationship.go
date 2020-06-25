package models

import (
	"Neo4jCURD/config"
	"fmt"
	"log"
)

type ResultRel struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    ReqRelationship `json:"data"`
}

type ReqRelationship struct {
	PersonId string  `json:"person_id"`
	MovieId  string  `json:"movie_id"`
	Rating   float32 `json:"rating"`
}

func AddRelationshipPersonMovie(req *ReqRelationship) (bool, error) {
	query := fmt.Sprintf(`MATCH	(p:Person), (m:Movie) 
								WHERE 	p.id = toInt(%s) 
								AND 	m.id = toInt(%s) 
								CREATE	(p)-[:RATED {rating: %f}]->(m) 
								RETURN p, m`, req.PersonId, req.MovieId, req.Rating)

	mapParams := map[string]interface{}{}

	DeleteRelPersonMovie(req.PersonId, req.MovieId)

	_, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return false, err
	}

	return true, err
}

// ???
func CheckRelPersonMovieExists(pid string, mid string) bool {
	// MATCH (p:Person {id: 99}), (m:Movie {id: 99}) RETURN EXISTS((p)-[:RATED]-(m))
	query := fmt.Sprintf(`MATCH (p:Person {id: toInt(%s)}), 
									(m:Movie {id: toInt(%s)})
								RETURN EXISTS((p)-[:RATED]-(m))`, pid, mid)

	mapPrams := map[string]interface{}{}

	result, _ := config.ResultQuery(query, mapPrams)
	log.Println(result.Record())
	if result.Record() != nil {
		log.Println(result.Record(), "nil")
		return true
	}

	return false
}

func DeleteRelPersonMovie(pid string, mid string) {
	// MATCH (p:Person {id: 99})-[r:RATED]->(m:Movie {id: 99}) DELETE r
	query := fmt.Sprintf(`MATCH (p:Person {id: toInt(%s)})
										-[r:RATED]->(m:Movie {id: toInt(%s)})
								DELETE r`, pid, mid)

	mapPrams := map[string]interface{}{}

	_, _ = config.ResultQuery(query, mapPrams)
}
