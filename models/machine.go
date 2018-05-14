package models

type machine struct{
	Id uint32 `gorm:"primary_key"`
	OwnerId uint32
	Url string
	CreateTime string
}

