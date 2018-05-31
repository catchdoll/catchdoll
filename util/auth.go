package util

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strings"
	"strconv"
	"github.com/catchdoll/conf"
)

func WxLogin(ctx *gin.Context) {
	code, success := ctx.GetQuery("code")
	if !success {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	//用code参数请求wx后台接口获取openid和sessionkey
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/sns/jscode2session?appid="+conf.GlobalConf.
		WxAppid+ "&secret=" + conf.GlobalConf.WxAppsecret + "&js_code=" + code + "&grant_type=authorization_code", nil)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var result WxLoginReturn
	fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	fmt.Println(result)
	fmt.Println(err)
	//return
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	//session_key进行一道加密，
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(result.SessionKey), bcrypt.DefaultCost)
	password := string(bytePassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	var user model.User
	model.DC.Debug().Where("openid = ? ", result.Openid).Find(&user)
	if user.Id == 0 {
		user.Openid = result.Openid
		user.Password = password
		model.DC.Debug().Create(&user)
		fmt.Println(user)
	} else {
		user.Password = password
		model.DC.Debug().Model(&user).Update("password", password)
	}
	trueExpireTime := result.Expire - 60
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Uid:      user.Id,
		Password: user.Password,
		Expire:   result.Expire,
	})
	//fmt.Println("ok")
	err = model.RC.Set("uid:"+strconv.Itoa(int(user.Id)), user.Password, time.Second*time.Duration(trueExpireTime)).Err()
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	//fmt.Println("ok")
	//fmt.Println(model.RC.Get("uid:2").Result())
	//fmt.Println("ok")
	//fmt.Println(trueExpireTime)
	//token := jwt.New(jwt.SigningMethodHS256)
	//claims := make(jwt.MapClaims)
	//claims["expire"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	//claims["expire"] = trueExpireTime
	//claims["uid"] = user.Id
	//claims["password"] = user.Password
	//token.Claims = claims
	key := "Jacku"
	tokenString, _ := token.SignedString([]byte(key))
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func WxAuth(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if splitToken[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		ctx.Abort()
		return
	}
	fmt.Println(splitToken)
	reqToken = splitToken[1]
	fmt.Println(reqToken)
	var user User
	token, err := jwt.ParseWithClaims(reqToken, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte("Jacku"), nil
	})
	if err != nil {
		//fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		ctx.Abort()
		return
	}
	if token.Valid && user.Uid != 0 {
		redisPwd, err := model.RC.Get("uid:" + strconv.Itoa(int(user.Uid))).Result()
		if err != nil || redisPwd != user.Password {
			//fmt.Println("error is", err)
			//fmt.Println("redis pwd is", redisPwd, "and user pwd is", user.Password)
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "mark3"})
			ctx.Abort()
			return
		}
		ctx.Set("uid",user.Uid)
		ctx.Next()
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "un"})
	ctx.Abort()
	return
}

type (
	WxLoginReturn struct {
		SessionKey string `json:"session_key"`
		Openid     string `json:"openid"`
		Expire     int    `json:"expire"`
	}
	WxLoginFailureReturn struct {
		ErrCode int
		ErrMsg  string
	}
	UserPwdPair struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	LoginResult struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
	}
	User struct {
		Uid      uint32 `json:"uid"`
		Password string `json:"password"`
		Expire   int    `json:"expire"`
		jwt.StandardClaims
	}
)

func GetUid(ctx *gin.Context)(uint32, bool){
	getUid := ctx.MustGet("uid")
	if uid, ok := getUid.(uint32); !ok{
		return 0, false
	}else{
		return uid, true
	}

}
