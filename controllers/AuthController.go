package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	// "strconv"
	// "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	// "github.com/dgrijalva/jwt-go"

	"yooyin/models"
	"yooyin/services"
)

type AuthController struct {
	BaseController
}

// Login 请求 POST 数据
type loginRequest struct {
	Code          string     `json:"code"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

// Login 请求 响应 数据
type loginResponse = wxUserInfo

// 向微信服务器请求 session_key 返回数据，谜之 openid unionid 大小写不同步
type wxResponse struct {
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

// 微信解密后的用户信息，谜之 openid unionid 大小写不同步
type wxUserInfo struct {
	OpenId    string    `json:"openId,omitempty"`
	UnionId   string    `json:"unionId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Language  string    `json:"language"`
}

func (this *AuthController) Login() {
	req := new(loginRequest)

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, req); err != nil {
		this.JsonResponse(1, "Bad Request", nil)
	}
	
	// 向微信服务器获取用户信息 openid 等
	userInfo := wxLogin(req)

	user := &models.User{
		OpenId: userInfo.OpenId,
		UnionId: userInfo.UnionId,
		NickName: userInfo.NickName,
		AvatarUrl: userInfo.AvatarUrl,
		Gender: userInfo.Gender,
		Country: userInfo.Country,
		Province: userInfo.Province,
		City: userInfo.City,
		Language: userInfo.Language,
	}
	services.UserService.LoginUser(user)

	this.SetSession("openId", user.OpenId)

	this.JsonResponse(0, "ok", user)
}

// 微信小程序登陆，返回解密后的用户信息
// @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/auth.code2Session.html
func wxLogin(req *loginRequest) *wxUserInfo {
	appid := beego.AppConfig.String("weixin::appid")
	secret := beego.AppConfig.String("weixin::appsecret")

	wxReq := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	wxReq.Param("grant_type", "authorization_code")
	wxReq.Param("js_code", req.Code)
	wxReq.Param("secret", secret)
	wxReq.Param("appid", appid)

	wxRes := new(wxResponse)
	wxReq.ToJSON(wxRes)
	
	if wxCheckSignature(req.Signature, req.RawData, wxRes.SessionKey) == false {
		return nil
	}

	wxFullUserinfo := wxDecryptUserinfo(wxRes.SessionKey, req.EncryptedData, req.IV)

	return wxFullUserinfo
}

// 检查签名
// @see https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#%E6%95%B0%E6%8D%AE%E7%AD%BE%E5%90%8D%E6%A0%A1%E9%AA%8C
func wxCheckSignature (signature string, rawData string, sessionKey string) bool {
	s := sha1.New()
	s.Write([]byte(rawData + sessionKey))
	sha1 := s.Sum(nil)
	sha1hash := hex.EncodeToString(sha1)

	return signature == sha1hash
}

// 解码用户数据
// @see https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E8%A7%A3%E5%AF%86%E7%AE%97%E6%B3%95
func wxDecryptUserinfo(sessionKey string, encryptedData string, iv string) *wxUserInfo {
	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := wxAesCBCDecrypt(ed, sk, i)

	if err != nil {
		return nil
	}

	wxuserinfo := new(wxUserInfo)
	err = json.Unmarshal(decryptedData, wxuserinfo)
	if err != nil {
		return nil
	}
	return wxuserinfo
}

func wxAesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}

	if len(encryptData) % blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptData))
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = wxPKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

func wxPKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
