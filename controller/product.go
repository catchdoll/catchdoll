package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/catchdoll/model"
)

type(
	ProductOnShelfParams struct{
		Name string `json:"name"`
		Seller uint32 `json:"seller"`
		ImgUrl string `json:"img_url"`

	}
)

func ProductOnShelf(ctx *gin.Context){
	name := ctx.PostForm("name")
	seller, err := strconv.Atoi(ctx.PostForm("seller"))
	imgUrl := ctx.PostForm("img_url")
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"illegal param img_url"})
	}
	price, err := strconv.ParseFloat(ctx.PostForm("price"),64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"illegal price"})
	}

	product := model.Product{Name:name,Seller:uint32(seller),Price:price,ImgUrl:imgUrl,Status:model.PRODUCT_INITIAL}
	model.DC.Debug().Create(&product)
	if product.Id != 0 {
		ctx.JSON(http.StatusOK,gin.H{"result":product,"message":"create product success"})
	}else{
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"create product failure, please try later"})
	}
}


func ProductOffShelf(ctx *gin.Context){
	seller := ctx.PostForm("seller")
	id := ctx.PostForm("id")
	var product model.Product
	model.DC.Where("id = ? and seller = ?", id, seller).Find(&product)
	if product.Id == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"can't find the product"})
	}
	model.DC.Delete(&product)
	if product.Id != 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"delete product failure"})
	}else{
		ctx.JSON(http.StatusOK,gin.H{"message":"delete product success"})
	}

}


func ProductEdit(ctx *gin.Context){


}


