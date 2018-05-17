package model

type VideoComment struct{
	Id uint32 `gorm:"primary_key"`
	Content string
	Uid uint32
	Status uint8
	CreateTime string
}
