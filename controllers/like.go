package controllers

import (
	"github.com/astaxie/beego"
	"yooyin/models"
)

type LikeController struct {
 	beego.Controller
}

type LikeReq struct {
	LikeType     bool   //是否喜欢对方
	UserId       string //主人态id
	TargetUserId string //目标用户id
}

type LikeRsp struct {
 IsMatched bool //是否匹配成功
}

func (this *LikeController) LikeService() {
	var likeRelationInformation models.LikeRelationInformation
	var likeList []models.LikeRelationInformation
	isMatched := false

	var rsp LikeRsp

	likeRelationInformation.UserId = this.GetString("UserId")
	likeRelationInformation.TargetUserId = this.GetString("TargetUserId")
	likeRelationInformation.IsLiked, _ = this.GetBool("LikeType")

	models.AddLikeRelationInformation(&likeRelationInformation)

	if likeRelationInformation.IsLiked == false {
		rsp.IsMatched = false
		JsonResult(&this.Controller, 0, "", rsp)
	}

	models.GetLikeRelastionInformationByUserId(&likeList, likeRelationInformation.TargetUserId)

	for _, v := range likeList {
		if v.TargetUserId == likeRelationInformation.UserId && v.IsLiked == true {
			isMatched = true
		}
	}

	rsp.IsMatched = isMatched

	JsonResult(&this.Controller, 0, "", rsp)
}
