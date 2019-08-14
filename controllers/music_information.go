package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"

	"yooyin/models"
)

// MusicInformationController operations for MusicInformation
type MusicInformationController struct {
	BaseController
}

type MusicInformationQueryResult struct {
	TotalCount       int64                     `json:"total_count"`
	MusicInformation []models.MusicInformation `json:"music_information"`
}

type IntertRequery struct {
	Type          int        `json:"type"`
	LikeFields    string     `json:"like_fields"`
}

// GetByNameAndType ...
// @Title Get by name and type
// @Description get MusicInformation
// @Param	name	query	string	true	"匹配的名称"
// @Param	type	query	integer	true	"匹配类型"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer，if not set, default is 10"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer, if not set, default is 0"
// @Success 200 {object} controllers.MusicInformationQueryResult
// @Failure 403
// @router /get_by_name_type [get]
func (c *MusicInformationController) GetByNameAndType() {
	var totalCount = int64(0)

	name := c.GetString("name")
	info_type, _ := c.GetInt("type")
	var limit int64 = 10
	var offset int64

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	musicInformations := []models.MusicInformation{}
	if cnt, err := models.GetMusicInformationsCountByNameType(name, info_type); err == nil {
		totalCount = cnt
	} else {
		JsonResult(&c.Controller, 1, err.Error(), nil)
	}

	err := models.GetMusicInformationsByNameType(&musicInformations, name, info_type, offset, limit)
	if err != nil {
		JsonResult(&c.Controller, 1, err.Error(), nil)
	}

	JsonResult(&c.Controller, 0, "", MusicInformationQueryResult{TotalCount: totalCount, MusicInformation: musicInformations})
}

// InsertUserLikeMusicInfo ...
// @Title Insert User Like Music
// @Description add user like music information
// @Param	uuid	query	string	false	"测试时需要填入的字段"
// @Param	type	query	integer	true	"匹配类型"
// @Param	like_fields	query	string	true	"用户喜欢的内容，json类型(json字符串)，具体由前端定义，后端入库不参与解析"
// @Success 200 {object} controllers.JsonReturnDataMessage
// @Failure 403
// @router /add_user_like_music [post]
func (c *MusicInformationController) InsertUserLikeMusicInfo() {
	req := new(IntertRequery)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.JsonResponse(1, "Bad Request", nil)
	}

	//uuid := c.GetString("uuid")
	uuid := c.GetSession("openId").(string)
	likeType := req.Type
	likeFields := req.LikeFields

	userLike := models.UserLikeMusicInfo{Uuid: uuid, Type: likeType, LikeFields: likeFields}
	_, err := models.AddUserLikeMusicInfo(&userLike)
	if err != nil {
		JsonResult(&c.Controller, 1, err.Error(), nil)
	}
	JsonResult(&c.Controller, 0, "success", nil)
}

// GetUserLikeMusicInfo ...
// @Title get User Like Music
// @Description get user like music information
// @Param	uuid	query	string	false	"openid"
// @Param	type	query	string	false	"类型"
// @Success 200 {object} []models.UserLikeMusicInfo
// @Failure 403
// @router /get_user_like_music [get]
func (c *MusicInformationController) GetUserLikeMusicInfo() {
	o := orm.NewOrm()
	uuid := c.GetString("uuid")
	likeType := c.GetString("type")
	if uuid == "" {
		uuid = c.GetSession("openId").(string)
	}

	query := o.QueryTable(new(models.UserLikeMusicInfo)).Filter("uuid", uuid)
	if likeType != "" {
		query = query.Filter("type", likeType)
	}

	var likeInfos []models.UserLikeMusicInfo
	_, err := query.OrderBy("-id").All(&likeInfos)
	if err != nil {
		JsonResult(&c.Controller, 1, err.Error(), nil)
	}
	JsonResult(&c.Controller, 0, "success", likeInfos)
}
