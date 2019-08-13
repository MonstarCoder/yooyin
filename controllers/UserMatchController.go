package controllers

import (
	"github.com/wxnacy/wgo/arrays"
	"math/rand"
	"yooyin/models"

)

type UserMatchController struct {
	BaseController
}

type MatchRequest struct {
	MatchType	int 	// 0=歌单匹配 1=音乐节目匹配 2=测试匹配
	//UserId 		string
}

type MatchResponse struct {
	MatchRate   float32
	CoLikeList  string
}

func RandInt64(min int64, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}



func (this *UserMatchController) GetUserMatch (req MatchRequest) (rsp []MatchResponse) {
	var MatchList []models.UserLikeMusicInfo
	matchType, _ := this.GetInt("MatchType")

	//下面接口覃良给，返回所有用户的列表
	models.GetAllUserLikeMusicInfo(&MatchList)

	//剔除已匹配用户
	var TargetList []string
	var LikeRelationList []models.LikeRelationInformation
	userid := this.GetSession("openId").(string)
	models.GetTargetUidById(&LikeRelationList, userid)
	for _, v := range LikeRelationList{
		TargetList = append(TargetList, v.TargetUserId)
	}
	for _, v  := range MatchList{
		index := arrays.ContainsString(TargetList, v.Uuid)
		if (v.Uuid != userid) && (index == -1) && (matchType == v.Type) {
			var tmpRsp MatchResponse
			tmpRsp.MatchRate = float32(RandInt64(8000, 9999)/100)
			tmpRsp.CoLikeList = v.LikeFields
			rsp = append(rsp, tmpRsp)
		}
	}
	//JsonResult(&this.Controller, 0, "", rsp)
	return rsp
}