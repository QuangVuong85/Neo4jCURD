package models

import (
	"Neo4jCURD/config"
	"fmt"
	"log"
)

var (
	MovieList map[string]*Movie
)

type Movies struct {
	Counts     int      `json:"counts"`
	ListMovies []*Movie `json:"movies"`
}

type Movie struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ReqMovie struct {
	Node  string `json:"node"`
	Label string `json:"label"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

// edit node movie
func editNodeMovie(movieId string, m *Movie) (string, error) {
	query := `
			MATCH (m:Movie)
			WHERE 
				m.id = toInt($movieid)
			SET
				m.name = $name`

	mapPrams := map[string]interface{}{
		"movieid": movieId,
		"name":    m.Name,
	}

	_, err := config.ResultQuery(query, mapPrams)
	if err != nil {
		return "false", err
	}

	return "true", nil
}

// delete node movie
func deleteNodeMovie(movieId string) (error) {
	query := fmt.Sprintf("MATCH (m:Movie) WHERE m.id = %s DETACH DELETE m",
		movieId)
	mapParams := map[string]interface{}{}

	_, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return err
	}

	return nil
}

// add node movie
func addNodeMovie(reqmovie *ReqMovie) (error) {
	query := fmt.Sprintf("CREATE (%s:%s {id: %s, name: \"%s\"})",
		reqmovie.Node, reqmovie.Label, reqmovie.Id, reqmovie.Name)

	mapParam := map[string]interface{}{}

	_, err := config.ResultQuery(query, mapParam)
	if err != nil {
		return err
	}

	return nil
}

// get node movie
func getNodeMovie(movieId string) (*Movie, error) {
	query := `
			MATCH (m:Movie)
			WHERE 
				m.id = toInt($movieid)
			RETURN
				m.id as id,
				m.name as name`

	mapParms := map[string]interface{}{
		"movieid": movieId,
	}

	result, err := config.ResultQuery(query, mapParms)
	if err != nil {
		return nil, err
	}

	var temp_Movie *Movie
	if result.Next() {
		id := fmt.Sprintf("%v", result.Record().GetByIndex(0))
		name := fmt.Sprintf("%v", result.Record().GetByIndex(1))

		temp_Movie = &Movie{
			id,
			name}
	}

	return temp_Movie, nil
}

// get all node movies
func getAllNodeMovies() map[string]*Movie {
	MovieList := make(map[string]*Movie)

	query := `
			MATCH (m:Movie)
			RETURN
				m.id as id,
				m.name as name`

	mapParams := map[string]interface{}{}
	result, err := config.ResultQuery(query, mapParams)
	if err != nil {
		log.Println("", err)
		return nil
	}

	for result.Next() {
		id := fmt.Sprintf("%v", result.Record().GetByIndex(0))
		name := fmt.Sprintf("%v", result.Record().GetByIndex(1))

		temp_Movie := &Movie{
			id,
			name}

		MovieList[id] = temp_Movie
	}

	return MovieList
}