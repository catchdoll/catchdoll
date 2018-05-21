package model

type MachineComment struct{
	Id uint32 `gorm:"primary_key",json:"Id"`
	MachineId uint32 `json:"machine_id"`
	Uid uint32 `json:"uid"`
	Score uint8 `json:"score"`
	Content string `json:"content"`
	CreateTime string `sql:"-" json:"create_time"`
}
