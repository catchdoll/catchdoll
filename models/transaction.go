package models

type Transaction struct{
	Id uint32 `gorm:"primary_key"`
	OrderId uint32
	payer uint32
	Receiver uint32
	Value float64
	Status uint8
	Type uint8
	CreateTime string
	Remark string
}
