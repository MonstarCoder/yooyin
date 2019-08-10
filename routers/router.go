// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"yooyin/controllers"


	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/music_information",
			beego.NSInclude(
				&controllers.MusicInformationController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(



	&controllers.LoginController{},
			),
		),
		// beego.Get('/', func(context *context.Context) {
		//    context.Output.Body([]byte("hello"))
		//})

	)
	beego.Router("/login", &controllers.LoginController{})
	beego.AddNamespace(ns)
}
