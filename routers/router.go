package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"yooyin/controllers"
)

func init() {
	nsAuth := beego.NewNamespace("/auth",
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
	)

	nsApi := beego.NewNamespace("/api/v1",
		beego.NSBefore(authFilter),

		beego.NSNamespace("/music_information",
			beego.NSInclude(
				&controllers.MusicInformationController{},
			),
		),

		beego.NSNamespace("/like",
			beego.NSInclude(
				&controllers.LikeController{},
			),
		),

		beego.NSRouter("/user_like_types", &controllers.UserLikeTypeController{}, "get:GetUserLikeTypes"),

		beego.NSRouter("/music_styles", &controllers.MusicStyleController{}, "get:GetMusicStyles"),

		beego.NSRouter("/user_match", &controllers.UserMatchController{}, "get:GetUserMatch"),
	)

	beego.AddNamespace(nsAuth)
	beego.AddNamespace(nsApi)
}

func authFilter(ctx *context.Context) {
	openid := ctx.Input.Session("openId")
	if openid == nil {
		ctx.Abort(403, "Forbidden")
	}
}
