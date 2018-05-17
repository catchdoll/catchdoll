package model

type(
	User struct{
		Id uint32 `gorm:"primary_key"`
		Username string
		Openid string
		Sex string
		ImgUrl string
		CreateTime string `sql:"-"`//自动生成
		Level uint8
		Cellphone uint32
	}
)
