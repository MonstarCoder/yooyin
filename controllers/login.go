package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"

	"yooyin/models"
)

//var globalSessions *session.Manager

//func init() {
//	sessionConfig := &session.ManagerConfig{
//		CookieName:"gosessionid",
//		EnableSetCookie: true,
//		Gclifetime:3600,
//		Maxlifetime: 3600,
//		Secure: true,
//		CookieLifeTime: 3600,
//		ProviderConfig: "./tmp",
//	}
//	globalSessions, _ = session.NewManager("memory",sessionConfig)
//	go globalSessions.GC()
//}

type LoginController struct {
	beego.Controller
}

type WXLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

type WXUserInfo struct {
	OpenID    string    `json:"openId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId,omitempty"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

type ResUserInfo struct {
	UserInfo      WXUserInfo `json:"userInfo"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

type LoginBody struct {
	Code     string          `json:"code"`
	UserInfo ResUserInfo     `json:"userInfo"`
}

func (this *LoginController) LoginByWeixin() {

	fmt.Println("begin LoginByWeixin")

	var lb LoginBody
	body := this.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &lb)
	//fmt.Print(alb)

	userInfo := Login(lb.Code, lb.UserInfo)
	if userInfo == nil {
		fmt.Println("userInfo==nil")
	}

	fmt.Println("login succeed")
	o := orm.NewOrm()

	var user models.YooyinUser
	usertable := new(models.YooyinUser)
	fmt.Println(usertable)
	fmt.Println("before query database")
	fmt.Println(userInfo.OpenID)
	//err = o.QueryTable(usertable.TableName()).Filter("weixin_openid", userInfo.OpenID).One(&user)
	err = o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	fmt.Println("after query database")
	if err == orm.ErrNoRows {
		newuser := models.YooyinUser{WeixinOpenid: userInfo.OpenID, Avatar: userInfo.AvatarUrl, Gender: userInfo.Gender,
			Nickname: userInfo.NickName}
		o.Insert(&newuser)
		o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	}

	userinfo := make(map[string]interface{})
	userinfo["nickname"] = user.Nickname
	userinfo["gender"] = user.Gender
	userinfo["avatar"] = user.Avatar
	userinfo["birthday"] = user.Birthday

	if _, err := o.Update(&user); err == nil {

	}

	// TODO user.Nickname should be a unique userid
	// TODO  save sessionIDs
	//sessionKey := Create(Int2String(user.Nickname))
	sessionKey := Create(user.Nickname)
	fmt.Println("sessionkey==" + sessionKey)

	rtnInfo := make(map[string]interface{})
	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo

	ReturnHTTPSuccess(&this.Controller, rtnInfo)
	this.ServeJSON()
}

var key = []byte("adfadf!@#2")
var expireTime = 20
func Create(userid string) string {

	claims := CustomClaims{
		userid, jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireTime)).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(key)

	if err == nil {
		return tokenstr
	}
	return ""
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

type CustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func ReturnHTTPSuccess(this *beego.Controller, val interface{}) {

	rtndata := HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   val,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
}

func Login(code string, fullUserInfo ResUserInfo) *WXUserInfo {

	fmt.Println("begin login")
	secret := beego.AppConfig.String("wx_appsecret")
	appid := beego.AppConfig.String("wx_appid")

	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/auth.code2Session.html
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	req.Param("grant_type", "authorization_code")
	req.Param("js_code", code)
	req.Param("secret", secret)
	req.Param("appid", appid)

	var res WXLoginResponse
	req.ToJSON(&res)

	fmt.Println("after httplib.Get")

	s := sha1.New()
	s.Write([]byte(fullUserInfo.RawData + res.SessionKey))
	sha1 := s.Sum(nil)
	sha1hash := hex.EncodeToString(sha1)

	fmt.Println(res.SessionKey)
	fmt.Println(fullUserInfo.RawData)
	fmt.Println(fullUserInfo.Signature)
	fmt.Println(sha1hash)

	if fullUserInfo.Signature != sha1hash {
		return nil
	}
	userinfo := DecryptUserInfoData(res.SessionKey, fullUserInfo.EncryptedData, fullUserInfo.IV)

	fmt.Println("return login")
	return userinfo

}

func DecryptUserInfoData(sessionKey string, encryptedData string, iv string) *WXUserInfo {

	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := AesCBCDecrypt(ed, sk, i)

	if err != nil {
		return nil
	}

	var wxuserinfo WXUserInfo
	//fmt.Println(string(decryptedData))
	err = json.Unmarshal(decryptedData, &wxuserinfo)
	if err != nil {

	}
	return &wxuserinfo
}

func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}

	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptData))
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = PKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}



