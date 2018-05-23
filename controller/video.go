package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/model"
	"net/http"
	"github.com/catchdoll/api"
	"strconv"
)

func VideoUpload(ctx *gin.Context){

}

func VideoDownload(ctx *gin.Context){

}

func VideoCommentCreate(ctx *gin.Context){

}

func VideoCommentIndex(ctx *gin.Context){

}

func VideoCommentShow(ctx *gin.Context){

}

func TopVideoIndex(ctx *gin.Context){
	var videos []model.Video
	model.DC.Where("top = 1").Find(&videos)
	if len(videos) != 0{
		ctx.JSON(http.StatusOK,gin.H{"status":api.OK,"result":videos})
	}else{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.RESULTNOTFOUND,"message":"sorry, can't find any top videos"})
	}
}

func VideoShow(ctx *gin.Context){
	video_id, err := strconv.Atoi(ctx.PostForm("video_id"))
	if err != nil || video_id == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{"status":api.PARAMITERILLEGAL,"message":"illegal filed video_id"})
	}
	var video model.Video
	model.DC.Find(&video)
	if video.Id == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"status":api.RESULTNOTFOUND,"message":"sorry, can't find any video"})
	}
	model.DC.Model(&video).Related(&video.VideoComments)
	ctx.JSON(http.StatusOK,gin.H{"status":api.OK, "result":video})
}
