package doll_machine

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"net/http"
)

func main(){
	router := gin.Default()
	r1 := router.Group("/product")
	{
		r1.POST("/productOn", productOnHandler)
		r1.POST("/productOff",productOffHandler)
	}
}

func init(){
	var err error
	dc, err = gorm.Open("")
	if err != nil{
		panic("can't connect to database")
	}
}

var dc *gorm.DB

type(
	User struct{
		Id uint32
		Username string
		Openid string
		Sex string
		ImgUrl string
		CreateTime string
		Level uint8
		Cellphone uint32
	}
	Product struct{
		Id uint32
		Name string
		Seller uint32
		Price float64
		Status uint8
		CreateTime string `sql:"-"`//自动生成
	}
)


func productOnHandler(ctx *gin.Context){
	name := ctx.PostForm("name")
	seller, err := strconv.Atoi(ctx.PostForm("seller"))
	if err != nil{

	}
	price, err := strconv.ParseFloat(ctx.PostForm("price"),64)
	if err != nil{

	}

	product := Product{Name:name,Seller:uint32(seller),Price:price}
	dc.Create(&product)
	if product.Id != 0 {
		ctx.JSON(http.StatusOK,gin.H{"result":product,"message":"create product success"})
	}else{
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"create product failure, please try later"})
	}

}

func productOffHandler(ctx *gin.Context){
	seller := ctx.PostForm("seller")
	id := ctx.PostForm("id")
	var product Product
	dc.Where("id = ? and seller = ?", id, seller).Find(&product)
	if product.Id == 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"can't find the product"})
	}
	dc.Delete(&product)
	if product.Id != 0{
		ctx.JSON(http.StatusNotFound,gin.H{"message":"delete product failure"})
	}else{
		ctx.JSON(http.StatusOK,gin.H{"message":"delete product success"})
	}

}
