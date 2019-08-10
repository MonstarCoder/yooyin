package routers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/context/param"
)

func init() {
  beego.GlobalControllerRouter["yooyin/controllers:MusicInformationController"] = append(
    beego.GlobalControllerRouter["yooyin/controllers:MusicInformationController"],
    beego.ControllerComments{
      Method: "GetByNameAndType",
      Router: `/get_by_name_type`,
      AllowHTTPMethods: []string{"get"},
      MethodParams: param.Make(),
      Filters: nil,
      Params: nil,
    },
  )
  beego.GlobalControllerRouter["yooyin/controllers:LikeController"] = append(
    beego.GlobalControllerRouter["yooyin/controllers:LikeController"],
    beego.ControllerComments{
      Method: "LikeService",
      Router: `/gogo`,
      AllowHTTPMethods: []string{"get"},
      MethodParams: param.Make(),
      Params: nil,
    },
  )
}
