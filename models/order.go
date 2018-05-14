package models

type Order struct{
	Id uint32 `gorm:"primary_key"`
	Buyer uint32
	Seller uint32
	Value float64
	Status uint8
	CreateTime string
	Remark string
}
