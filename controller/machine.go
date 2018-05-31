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

type(
	MachineCreateParams struct{
		Address string `json:"address"`
		Lat string `json:"lat"`
		Lon string `json:"lon"`
		Url string `json:"url"`
	}
	MachineCommentCreateParams struct{
		Content string `json:"content"`
		MachineId uint32 `json:"machine_id"`
		Score uint8 `json:"score"`
	}
)

func TopMachinesIndex(ctx *gin.Context){//查找所有娃娃机信息
	var machines []model.Machine
	model.DC.Where("top = 1").Find(&machines)
	if len(machines) == 0{
		ctx.JSON(http.StatusOK,gin.H{"status":api.RESULTNOTFOUND,"message":"no top machines found"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machines})
	return
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
	if machine.Id == 0{
		ctx.JSON(http.StatusOK, gin.H{"status":api.RESULTNOTFOUND,"message":"can't find machine"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machine})
	return
}

func MachineCommentCreate(ctx *gin.Context){
	var input MachineCommentCreateParams
	if ctx.BindJSON(&input) != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"plz check the params format"})
		return
	}
	//paramUid, err := strconv.Atoi(ctx.PostForm("uid"))
	//fmt.Println(ctx.PostForm("uid"))
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal parameter uid_id"})
	//	return
	//}
	//uid := uint32(paramUid)
	//查找有没有这台机器
	uid ,ok := util.GetUid(ctx)
	if input.MachineId == 0 || input.Score > 5 || !ok{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"plz check the params format"})
		return
	}
	var machine model.Machine
	model.DC.Where("Id = ?", input.MachineId).Find(&machine)
	if machine.Id == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"can't find machine with specified machine_id"})
		return
	}
	var machineComment model.MachineComment
	model.DC.Where("machine_id = ? and uid = ?",input.MachineId, uid).Find(&machineComment)
	if machineComment.Id != 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"you have commented, do not repeat"})
		return
	}
	machineComment = model.MachineComment{
		Content:input.Content,
		MachineId:input.MachineId,
		Score:input.Score,
		Uid:uid,
	}
	model.DC.Create(&machineComment)
	if machineComment.Id == 0{
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":api.DATAORPERATIONFAILURE, "message":"machine creation failure"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machineComment})

}

func MachineCreate(ctx *gin.Context){
	var input MachineCreateParams
	if ctx.BindJSON(&input) != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"plz check the params format"})
		return
	}
	uid, ok := util.GetUid(ctx)
	fmt.Println("uid is",uid)
	if input.Address == "" || input.Url == "" ||  !ok{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL, "message":"plz check the params format"})
		return
	}
	newMachine := model.Machine{
		Url:input.Url,
		Address:input.Address,
		Lat:input.Lat,
		Lon:input.Lon,
		CreatorId:uid,
	}
	model.DC.Debug().Create(&newMachine)
	if newMachine.Id == 0{
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":api.DATAORPERATIONFAILURE, "message":"create machine failure"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":newMachine})
}


