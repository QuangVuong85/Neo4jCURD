package controllers

import (
	"Neo4jCURD/models"
	"fmt"
	"github.com/astaxie/beego"
)

// Movie Recommendations
type QueryController struct {
	beego.Controller
}

// @Title Get
// @Description get all Movie Recommendations for `Name Person` ex: Michael Hunger
// @Param	personName		query 	string	true		"The key for staticblock"
// @Success 200 {object} models.RespMoviesRec
// @Failure 403 :personName not exists
// @router /GetDatax [get]
func (p *QueryController) Get() {
	var personName string
	err := p.Ctx.Input.Bind(&personName,"personName")

	if err != nil {
		fmt.Println("err = ",err)
	}
	results, _ := models.GetMovieRecommendationsPerson(personName)
	/*list := make([]*models.MovieRec, 0)

	for _, i := range results {
		list = append(list, i)
	}*/

	var message string
	if len(results) == 0 {
		message = "Person have name " + personName + " not in Nodes Person."
	} else {
		message = "Get Movie Recommendations for " + personName
	}

	p.Data["json"] = models.RespMoviesRec{message, results}
	p.ServeJSON()
}

// @Title Add Cosine Similarities to the Graph
// @Description add Cosine Similarities to the Graph
// @Success 200 {Status} is true
// @Failure 403 {Status} is false
// @router / [post]
func (q *QueryController) Post() {
	err := models.AddCosineSimilarities()

	var ok string
	if err != nil {
		ok = "false"
	} else {
		ok = "true"
	}

	q.Data["json"] = map[string]string{"Status": ok}
	q.ServeJSON()
}
