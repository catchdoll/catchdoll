package model

type Product struct{
	Id uint32 `gorm:"primary_key"`
	Name string
	Seller uint32
	Price float64
	Status uint8
	CreateTime string `sql:"-"`//自动生成
	ImgUrl string
}

const (
	PRODUCT_INITIAL = 0
	PRODUCT_ON_SHELF = 1
	PRODUCT_SOLD_OUT = 2
	PRODUCT_OFF_SHELF_CHECK = 3
	PRODUCT_OFF_SHELF = 4
)
