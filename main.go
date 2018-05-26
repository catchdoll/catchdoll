package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/catchdoll/model"
	"github.com/catchdoll/controller"
	"time"
	"github.com/appleboy/gin-jwt"
	"net/http"
	"github.com/catchdoll/util"
)

func main() {
	AuthMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "catchdoll",
		Key:        []byte("catchdollsecrect"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: util.Authenticate,
		Authorizator: util.Authorize,
		Unauthorized: func (c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//PayloadFunc:util.Payload,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	if model.InitDB() != nil {
		panic("can't connect to database")
	}
	//var err error
	//DC, err = gorm.Open("mysql","root:194466@/doll_machine?charset=utf8")
	//if err != nil{
	//	panic("db connection failure")
	//}
	router := gin.Default()
	router.POST("/login",AuthMiddleware.LoginHandler)
	r1 := router.Group("/")
	r1.Use(AuthMiddleware.MiddlewareFunc())
	{
		r1.POST("/hello",productHandler)
		r1.POST("/productOn", controller.ProductOnShelf)
		r1.POST("/productOff", controller.ProductOffShelf)
		r1.GET("/machine_top/", controller.TopMachinesIndex)
		r1.GET("/machine/:id", controller.MachineShow)
		r1.POST("/machine_comment", controller.MachineCommentCreate)

	}
	router.Run(":8000")

}

func productHandler(c *gin.Context){
	key := c.PostForm("key")
	c.JSON(http.StatusOK,gin.H{"result":key})
}

//var DC *gorm.DC


