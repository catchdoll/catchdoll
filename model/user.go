package model

type(
	User struct{
		Id uint32 `sql:"primary_key"`
		Username string
		Openid string
		Sex string `gorm:"default:'M'"`
		ImgUrl string
		CreateTime string `sql:"-"`//自动生成
		Level uint8
		Cellphone uint32
		Password string
	}
)
