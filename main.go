package main

import (
	"github.com/gin-gonic/gin"
	"github.com/catchdoll/models"
	"github.com/catchdoll/controllers"
)

func main(){

	if models.InitDB() != nil{
		panic("can't connect to database");
	}

	router := gin.Default()
	r1 := router.Group("/product")
	{
		r1.POST("/productOn",controllers.ProductOn)
		r1.POST("/productOff",controllers.ProductOff)
	}
}






