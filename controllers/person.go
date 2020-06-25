package controllers

import (
	"Neo4jCURD/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Person
type PersonController struct {
	beego.Controller
}

// @Title Get
// @Description get person by personId
// @Param	personId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Person
// @Failure 403 :personId is empty
// @router /:personId [get]
func (p *PersonController) Get() {
	personId := p.GetString(":personId")

	if personId != "" {
		person, err := models.GetPerson(personId)

		if err != nil {
			p.Data["json"] = err.Error()
		} else {
			listPerson := make([]*models.Person, 0)
			listPerson = append(listPerson, person)
			p.Data["json"] = models.Persons{len(listPerson), listPerson}
		}
	}

	p.ServeJSON()
}

// @Title GetAll
// @Description get all Persons
// @Success 200 {object} models.Person
// @router / [get]
func (p *PersonController) GetAll() {
	persons := models.GetAllPersons()
	listPerson := make([]*models.Person, 0)

	for _, i := range persons {
		listPerson = append(listPerson, i)
	}

	p.Data["json"] = models.Persons{len(listPerson), listPerson}
	p.ServeJSON()
}

// @Title Update
// @Description update the person
// @Param	personId		path 	string	true		"The personId you want to update"
// @Param	body		body 	models.Person	true		"body for person content"
// @Success 200 {object} models.Person
// @Failure 403 :personId is not int
// @router /:personId [put]
func (p *PersonController) Put() {
	personId := p.GetString(":personId")

	if personId != "" {
		var person models.Person
		json.Unmarshal(p.Ctx.Input.RequestBody, &person)

		mo, err := models.UpdatePerson(personId, &person)

		if err != nil {
			p.Data["json"] = err.Error()
		} else {
			p.Data["json"] = mo
		}
	}

	p.ServeJSON()
}

// @Title CreatePerson
// @Description create person
// @Param	body		body 	models.ReqPerson	true		"body for person content"
// @Success 200 {int} models.Person.Id
// @Failure 403 body is empty
// @router / [post]
func (p *PersonController) Post() {
	var reqperson models.ReqPerson

	json.Unmarshal(p.Ctx.Input.RequestBody, &reqperson)

	personId := models.AddPerson(&reqperson)

	p.Data["json"] = map[string]string{"personId": personId}
	p.ServeJSON()
}

// @Title Delete
// @Description delete the person
// @Param	personId		path 	string	true		"The personId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 personId is empty
// @router /:personId [delete]
func (p *PersonController) Delete() {
	personId := p.GetString(":personId")
	err := models.DeletePerson(personId)
	p.Data["json"] = err
	p.ServeJSON()
}
