package models

import (
	"Neo4jCURD/config"
	"fmt"
)

type MovieRec struct {
	Movie          string `json:"movie"`
	Recommendation string `json:"recommendation"`
}

type RespMoviesRec struct {
	Messages string      `json:"messages"`
	Data     []*MovieRec `json:"data"`
}

// Get Movie Recommendations for Person(name=?) // Michael Hunger
func GetMovieRecommendationsPerson(namePerson string) ([]*MovieRec, error) {
	//fmt.Println(strings.Replace()namePerson)

	//results := make(map[string]*MovieRec)
	query := `MATCH    (b:Person)-[r:RATED]->(m:Movie), (b)-[s:SIMILARITY]-(a:Person {name: $nameperson})
				WHERE    NOT((a)-[:RATED]->(m))
				WITH     m, s.similarity AS similarity, r.rating AS rating
				ORDER BY m.name, similarity DESC
				WITH     m.name AS movie, COLLECT(rating)[0..3] AS ratings
				WITH     movie, REDUCE(s = 0, i IN ratings | s + i)*1.0 / LENGTH(ratings) AS reco
				ORDER BY reco DESC
				RETURN   movie AS Movie, reco AS Recommendation`

	mapParams := map[string]interface{}{
		"nameperson": namePerson,
	}

	result, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return nil, err
	}

	ListMovie := []*MovieRec{}
	//i := 0
	for result.Next() {
		movie := fmt.Sprintf("%v", result.Record().GetByIndex(0))
		recommendation := fmt.Sprintf("%v", result.Record().GetByIndex(1))

		temp_MovieRec := &MovieRec{
			movie,
			recommendation,
		}

		ListMovie = append(ListMovie, temp_MovieRec)
		//results[string(i)] = temp_MovieRec
		//i++
	}

	return ListMovie, nil
}

// Add Cosine Similarities to the Graph
func AddCosineSimilarities() error {
	query := `	MATCH (p1:Person)-[x:RATED]->(m:Movie)<-[y:RATED]-(p2:Person)
				WITH    SUM(x.rating * y.rating) AS xyDotProduct,
						SQRT(REDUCE(xDot = 0.0, a IN COLLECT(x.rating) | xDot + a^2)) AS xLength,
						SQRT(REDUCE(yDot = 0.0, b IN COLLECT(y.rating) | yDot + b^2)) AS yLength,
						p1, p2
				MERGE (p1)-[s:SIMILARITY]-(p2)
				SET s.similarity = xyDotProduct / (xLength * yLength)`

	mapParams := map[string]interface{}{}

	_, err := config.ResultQuery(query, mapParams)
	if err != nil {
		return err
	}

	return nil
}
