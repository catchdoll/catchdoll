package models

import "github.com/jinzhu/gorm"

func InitDB() (error) {
	var err error
	DB, err = gorm.Open("")
	return err
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	//if err == nil {
	//	DB = db
		//db.LogMode(true)
		//db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")

	//}

}

var DB *gorm.DB
