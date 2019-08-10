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
	var contactInformation models.ContactInformation
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

	if isMatched == true {
		contactInformation.UserId = this.GetString("UserId")
		contactInformation.ContactUserId = this.GetString("TargetUserId")
		models.AddContactInformation(&contactInformation)

		contactInformation.UserId = this.GetString("TargetUserId")
		contactInformation.ContactUserId = this.GetString("UserId")
		models.AddContactInformation(&contactInformation)
	}

	JsonResult(&this.Controller, 0, "", rsp)
}
