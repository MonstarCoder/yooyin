package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["yin_you/controllers:MusicInformationController"] = append(beego.GlobalControllerRouter["yin_you/controllers:MusicInformationController"],
		beego.ControllerComments{
			Method: "GetByNameAndType",
			Router: `/get_by_name_type`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
