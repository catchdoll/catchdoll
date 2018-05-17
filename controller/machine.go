package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/catchdoll/api"
	"github.com/catchdoll/model"
)

func MachineCreate(ctx *gin.Context){//新增娃娃机信息
	//todo 完善管理员身份的验证
	owner, err := strconv.Atoi(ctx.PostForm("owner"))
	if err != nil || owner == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.PARAMITERILLEGAL,"message":"owner illegal"})
	}
	address := ctx.PostForm("address")
	if address == ""{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.PARAMITERILLEGAL,"message":"address illegal"})
	}
	coordinate := ctx.PostForm("coordinate")
	if coordinate == ""{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.PARAMITERILLEGAL,"message":"coordinate illegal"})
	}
	machine := model.Machine{Owner:uint32(owner),Address:address,Coordinate:coordinate}
	model.DC.Create(&machine)
}

func MachineDelete(ctx *gin.Context){

}

func MachineEdit(ctx *gin.Context){

}

func MachineIndex(ctx *gin.Context){

}

func MachineShow(ctx *gin.Context){

}

func MachineCommentCreate(ctx *gin.Context){

}

func MachineCommentIndex(ctx *gin.Context){

}

func MachineCommentShow(ctx *gin.Context){

}

func MachineCommentDelete(ctx *gin.Context){

}



