package models

type MachineComment struct{
	Id uint32 `gorm:"primary_key"`
	MachineId uint32
	Uid uint32
	Score uint8
	Content string
	CreateTime string
}
