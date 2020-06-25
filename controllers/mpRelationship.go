package controllers

import (
	"Neo4jCURD/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Relationship 2 nodes Person and Movie
type RelMoviePerson struct {
	beego.Controller
}

// @Title Relationship 2 nodes Person and Movie
// @Description create ReqRelationship
// @Param	body		body 	models.ReqRelationship	true		"body for ReqRelationship content"
// @Success 200 {ReqRelationship} models.ReqRelationship
// @Failure 403 body is empty
// @router / [post]
func (rel *RelMoviePerson) Post() {
	var reqRelationship models.ReqRelationship

	json.Unmarshal(rel.Ctx.Input.RequestBody, &reqRelationship)

	status, _ := models.AddRelationshipPersonMovie(&reqRelationship)

	rel.Data["json"] = models.ResultRel{ status, "",reqRelationship}
	rel.ServeJSON()
}
