package controllers

import (
	"Neo4jCURD/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Movie
type MovieController struct {
	beego.Controller
}

// @Title Get
// @Description get user by movieId
// @Param	movieId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Movie
// @Failure 403 :movieId is empty
// @router /:movieId [get]
func (m *MovieController) Get() {
	movieId := m.GetString(":movieId")
	if movieId != "" {
		movie, err := models.GetMovie(movieId)

		if err != nil {
			m.Data["json"] = err.Error()
		} else {
			listMovies := make([]*models.Movie, 0)
			listMovies = append(listMovies, movie)
			m.Data["json"] = models.Movies{len(listMovies), listMovies}
		}
	}

	m.ServeJSON()
}

// @Title GetAll
// @Description get all Movies
// @Success 200 {object} models.Movies
// @router / [get]
func (m *MovieController) GetAll() {
	movies := models.GetAllMovies()

	listMovies := make([]*models.Movie, 0)
	for _, i := range movies {
		listMovies = append(listMovies, i)
	}

	m.Data["json"] = models.Movies{len(listMovies), listMovies}
	m.ServeJSON()
}

// @Title CreateMovie
// @Description create movie
// @Param	body		body 	models.ReqMovie	true		"body for movie content"
// @Success 200 {int} models.Movie.movieId
// @Failure 403 body is empty
// @router / [post]
func (m *MovieController) Post() {
	var reqmovie models.ReqMovie

	json.Unmarshal(m.Ctx.Input.RequestBody, &reqmovie)

	movieId := models.AddMovie(&reqmovie)

	m.Data["json"] = map[string]string{"movieId": movieId}
	m.ServeJSON()
}

// @Title Update
// @Description update the movie
// @Param	movieId		path 	string	true		"The movieId you want to update"
// @Param	body		body 	models.Movie	true		"body for movie content"
// @Success 200 {object} models.Movie
// @Failure 403 :movieId is not int
// @router /:movieId [put]
func (m *MovieController) Put() {
	movieId := m.GetString(":movieId")

	if movieId != "" {
		var movie models.Movie
		json.Unmarshal(m.Ctx.Input.RequestBody, &movie)

		mo, err := models.UpdateMovie(movieId, &movie)

		if err != nil {
			m.Data["json"] = err.Error()
		} else {
			m.Data["json"] = mo
		}
	}

	m.ServeJSON()
}

// @Title Delete
// @Description delete the movie
// @Param	movieId		path 	string	true		"The movieId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 movieId is empty
// @router /:movieId [delete]
func (m *MovieController) Delete() {
	movieId := m.GetString(":movieId")
	err := models.DeleteMovie(movieId)
	m.Data["json"] = err
	m.ServeJSON()
}
