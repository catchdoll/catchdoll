package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/catchdoll/models"
)

func ProductOn(ctx *gin.Context){
	name := ctx.PostForm("name")
	seller, err := strconv.Atoi(ctx.PostForm("seller"))
	if err != nil{

	}
	price, err := strconv.ParseFloat(ctx.PostForm("price"),64)
	if err != nil{

	}

	product := models.Product{Name:name,Seller:uint32(seller),Price:price}
	models.DB.Create(&product)
	if product.Id != 0 {
		ctx.JSON(http.StatusOK,gin.H{"result":product,"message":"create product success"})
	}else{
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"create product failure, please try later"})
	}
}

func ProductOff(ctx *gin.Context){
	seller := ctx.PostForm("seller")
	id := ctx.PostForm("id")
	var product models.Product
	models.DB.Where("id = ? and seller = ?", id, seller).Find(&product)
	if product.Id == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"can't find the product"})
	}
	models.DB.Delete(&product)
	if product.Id != 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"delete product failure"})
	}else{
		ctx.JSON(http.StatusOK,gin.H{"message":"delete product success"})
	}

}


func ProductEdit(ctx *gin.Context){

}
