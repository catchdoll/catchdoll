package models

type Product struct{
	Id uint32 `gorm:"primary_key"`
	Name string
	Seller uint32
	Price float64
	Status uint8
	CreateTime string `sql:"-"`//自动生成
	ImgUrl string
}
