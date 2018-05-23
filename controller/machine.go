package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
	"strconv"
	"net/http"
	"github.com/catchdoll/api"
	"fmt"
	"github.com/catchdoll/util"
)


func TopMachinesIndex(ctx *gin.Context){//查找所有娃娃机信息
	var machines []model.Machine
	model.DC.Where("top = 1").Find(&machines)
	if len(machines) == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.RESULTNOTFOUND,"message":"no top machines found"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machines})
}

func MachineShow(ctx *gin.Context){//展示单个娃娃机的信息(包含所有评论)
	paramMachineId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || paramMachineId == 0 {
		fmt.Println("id is",paramMachineId)
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal parameter machine_id"})
		return
	}
	machineId := paramMachineId
	var machine model.Machine
	model.DC.Where("id = ?", machineId).Find(&machine)
	model.DC.Debug().Model(&machine).Related(&machine.MachineComments)
	if machine.Id == 0{
		ctx.JSON(http.StatusOK, gin.H{"status":api.RESULTNOTFOUND,"message":"can't find machine"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machine})
	return
}

func MachineCommentCreate(ctx *gin.Context){
	content := ctx.PostForm("content")
	paramScore, err := strconv.Atoi(ctx.PostForm("score"))
	if err != nil || paramScore > 5 || paramScore < 1{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal parameter score"})
		return
	}
	score := uint8(paramScore)
	paramMachineId, err := strconv.Atoi(ctx.PostForm("machine_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal parameter machine_id"})
		return
	}
	machineId := uint32(paramMachineId)
	//paramUid, err := strconv.Atoi(ctx.PostForm("uid"))
	//fmt.Println(ctx.PostForm("uid"))
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal parameter uid_id"})
	//	return
	//}
	//uid := uint32(paramUid)
	uid, err := util.GetUid(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized,gin.H{"status":api.UNAUTHORIZED,"message":"authorization problem"})
	}
	//查找有没有这台机器
	var machine model.Machine
	model.DC.Where("Id = ?", machineId).Find(&machine)
	if machine.Id == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"can't find machine with specified machine_id"})
		return
	}
	var machineComment model.MachineComment
	model.DC.Where("machine_id = ? and uid = ?",machineId, uid).Find(&machineComment)
	if machineComment.Id != 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"you have commented, do not repeat"})
		return
	}
	machineComment = model.MachineComment{
		Content:content,
		MachineId:machineId,
		Score:score,
		Uid:uid,
	}
	fmt.Println(score)
	model.DC.Create(&machineComment)
	if machineComment.Id == 0{
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":api.DATAORPERATIONFAILURE, "message":"machine creation failure"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machineComment})

}

func MachineCommentIndex(ctx *gin.Context){

}

func MachineCommentShow(ctx *gin.Context){

}

func MachineCommentDelete(ctx *gin.Context){

}

func MachineTest(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{"result":"ok"})
}



