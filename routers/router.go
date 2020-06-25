// @APIVersion 1.0.0
// @Title API DOCUMENTATION
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact quangvuong0805@gmail.com
// @TermsOfServiceUrl http://vuongdq.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Neo4jCURD/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/movie",
			beego.NSInclude(
				&controllers.MovieController{},
			),
		),
		beego.NSNamespace("/person",
			beego.NSInclude(
				&controllers.PersonController{},
			),
		),
		beego.NSNamespace("/relpersonmovie",
			beego.NSInclude(
				&controllers.RelMoviePerson{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
