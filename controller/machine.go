package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
	"strconv"
	"net/http"
	"github.com/catchdoll/api"
	"fmt"
	"github.com/catchdoll/util"
	"io/ioutil"
	"net/url"
)

type(
	MachineCreateParams struct{
		Address string `json:"address"`
		Lat string `json:"lat"`
		Lon string `json:"lon"`
		//Url string `json:"url"`
	}
	MachineCommentCreateParams struct{
		Content string `json:"content"`
		MachineId uint32 `json:"machine_id"`
		Score uint8 `json:"score"`
	}
	MachineCommentIndexParams struct{
		MachineId uint32 `json:"machine_id"`
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
	if input.Address == ""  ||  !ok{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL, "message":"plz check the params format"})
		return
	}

	ts := model.DC.Begin()
	lbsId, err := util.CreateLbsAddress(input.Address,input.Lat, input.Lon)
	if err != nil{
		ts.Rollback()
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":api.DATAORPERATIONFAILURE, "message":"create machine failure"})
		return
	}
	newMachine := model.Machine{
		//Url:input.Url,
		Address:input.Address,
		Lat:input.Lat,
		Lon:input.Lon,
		CreatorId:uid,
		LbsId:lbsId,
	}
	ts.Debug().Create(&newMachine)
	if newMachine.Id == 0{
		ts.Rollback()
		//todo lbs表未创建成功本地表却无法创建，必须记录下来，作异常处理
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":api.DATAORPERATIONFAILURE, "message":"create machine failure"})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":newMachine})
}

func MachineCommentIndex(ctx *gin.Context){
	paramMachineId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || paramMachineId == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL, "message":"illegal param id"})
		return
	}
	var machine model.Machine
	var machineComments []model.MachineComment
	model.DC.Where("id = ?", paramMachineId).Find(&machine)
	if machine.Id == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL, "message":"can't find machine by id"})
		return
	}
	model.DC.Model(&machine).Related(machineComments)
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"message":"success","result":machineComments})
}

func MachineCollect(ctx *gin.Context){//收藏娃娃机
	machineId := ctx.MustGet("id")
	if machineId == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal param id"})
		return
	}
	var machineCollectRecord model.MachineCollect
	model.DC.Where("id = ?", machineId).Find(&machineCollectRecord)
	if machineCollectRecord.Id != 0{
		ctx.JSON(http.StatusOK,gin.H{"status":api.DATAORPERATIONFAILURE,"message":"collect machine failure"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"stattus":api.OK,"message":"success"})
}


func MachineShare(ctx *gin.Context){//分享娃娃机

}

func MachineVote(ctx *gin.Context){//给娃娃机点赞

}

func LbsTest(ctx *gin.Context){
	uri := "/geodata/v4/poi/create"
	form := url.Values{
		"ak":{"6wVrtUncrfIq5SE4AlYtrCRZtDtsT1kP"},
		//"id":{"190008"},
		//"name":{"doll_machine_address"},
		"coord_type":{"3"},
		"geotable_id":{"1000003990"},
		"latitude":{"22.56124"},
		"longitude":{"114.106875"},
		"title":{"复制"},
	}
	SNCode, err := util.GetLbsSN(uri,form)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"error get ak"})
	}
	form.Add("sn",SNCode)
	resp, err := http.PostForm("http://api.map.baidu.com/geodata/v4/poi/create",form)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"request failed"})
	}
	result, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}






