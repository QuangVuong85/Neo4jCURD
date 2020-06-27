// @APIVersion 1.0.0
// @Title API DOCUMENTATION
// @Description Beego API with JWT implementation and Neo4j database
// @Contact quangvuong0805@gmail.com
// @TermsOfServiceUrl http://vuongdq.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Neo4jCURD/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.Get("/", func(ctx *context.Context) {
		_ = ctx.Output.Body([]byte("This is a Beego + JWT API - Creator: vuongdq85 (quangvuong0805@gmail.com)"))
	})

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
		beego.NSNamespace("/query",
			beego.NSInclude(
				&controllers.QueryController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
