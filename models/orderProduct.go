package models

type OrderProduct struct{
	Id uint32 `gorm:"primary_key"`
	OrderId uint32
	ProductId uint32
	ProductName string
	Quantity uint8
	CreateTime string
	Snapshot string
}
