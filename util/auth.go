package util

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
	"github.com/appleboy/gin-jwt"
	"errors"
	"strconv"
)

func Authenticate(userId string, password string, c *gin.Context) (string, bool) {
	var user model.User
	model.DC.Debug().Where("id = ? and password = ?",userId,password).Find(&user)
	return userId, user.Id != 0
}

func Authorize(userId string, c *gin.Context) bool {
	return true
}

type UserFetcher interface {
	getUid() (uint32, error)
}

func GetUid(ctx *gin.Context)(uint32, error){
	payload := jwt.ExtractClaims(ctx)
	strUid , success := payload["id"].(string)
	if !success{
		return 0, errors.New("payload id type is not a string")
	}
	uid, err := strconv.Atoi(strUid)
	if err != nil{
		return 0, err
	}
	return uint32(uid),nil
}






