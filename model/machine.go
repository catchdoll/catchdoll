package model

type Machine struct{
	Id uint32 `gorm:"primary_key"`
	Owner uint32
	Url string
	CreateTime string
	Address string
	Coordinate string
	Creator uint32
}

