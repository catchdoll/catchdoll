package model

type MachineShare struct{
	Id uint32 `gorm:"primary_key"`
	MachineId uint32
	Uid uint32
	CreateTime string `sql:"-"`
}
