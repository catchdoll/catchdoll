package util

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
)

func Authenticate(userId string, password string, c *gin.Context) (string, bool) {
	var user model.User
	model.DC.Debug().Where("username = ? and password = ?",userId,password).Find(&user)
	return userId, user.Id == 0
}

func Authorize(userId string, c *gin.Context) bool {
	return true
}





