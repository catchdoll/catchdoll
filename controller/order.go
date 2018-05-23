package controller

import (
	"github.com/gin-gonic/gin"
)

func OrderCreate(ctx *gin.Context){
	productId := ctx.PostForm("product")
}

