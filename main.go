package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/catchdoll/model"
	"github.com/catchdoll/controller"
	"github.com/catchdoll/util"
	"github.com/catchdoll/conf"
)

func main() {
	if conf.InitConfig() != nil{
		panic("can't initial configuration")
	}
	if model.InitDB() != nil {
		panic("can't connect to database")
	}
	if !model.InitRedis(){
		panic("can't connect to redis")
	}
	//var err error
	//DC, err = gorm.Open("mysql","root:194466@/doll_machine?charset=utf8")
	//if err != nil{
	//	panic("db connection failure")
	//}
	router := gin.Default()
	//router.POST("/login",AuthMiddleware.LoginHandler)
	router.GET("/wxLogin",util.WxLogin)
	r1 := router.Group("/")
	r1.Use(util.WxAuth)
	{
		//r1.POST("/productOn", controller.ProductOnShelf)//商品上架
		//r1.POST("/productOff", controller.ProductOffShelf)//商品下架
		r1.GET("/machine_top", controller.TopMachinesIndex)//置顶娃娃机信息
		r1.GET("/machine/:id", controller.MachineShow)//单个娃娃机信息(含评论)
		r1.POST("/machine_comment", controller.MachineCommentCreate)//评论娃娃机
		//r1.GET("/video_top/", controller.TopVideosIndex)
		//r1.GET("/video/:id",controller.VideoShow)
		r1.POST("/machine",controller.MachineCreate)

	}
	router.Run(":"+conf.GlobalConf.ServerPort)

}


//var DC *gorm.DC


