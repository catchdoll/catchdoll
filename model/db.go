package model

import "github.com/jinzhu/gorm"


func InitDB() (error) {
	var err error
	DC, err = gorm.Open("mysql","root:194466@/doll_machine?charset=utf8")
	return err
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	//if err == nil {
	//	DC = db
		//db.LogMode(true)
		//db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")

	//}

}

var DC *gorm.DB
