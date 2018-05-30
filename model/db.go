package model

import (
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
	"github.com/catchdoll/conf"
)


func InitDB() (error) {
	var err error
	//DC, err = gorm.Open("mysql", "catchdoll:+HNbtdBY@tcp(47.98.55.170:3306)/catchdoll?charset=utf8")
	//DC, err = gorm.Open("mysql","root:194466@/doll_machine?charset=utf8")
	DC, err = gorm.Open("mysql", conf.GlobalConf.Dsn)
	return err
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	//if err == nil {
	//	DC = db
		//db.LogMode(true)
		//db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")

	//}

}

func InitRedis()bool{
	RC = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:conf.GlobalConf.RedisAddr,
		Password: conf.GlobalConf.RedisPwd,
		DB:conf.GlobalConf.RedisDb,
	})
	return RC != nil
}

const (
	APPID = "wx09d7722244d6dc84"
	APPSECRET = "61a701a90198df394274fa68b00eda41"
)

var DC *gorm.DB

var RC *redis.Client
