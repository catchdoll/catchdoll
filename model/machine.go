package model

type Machine struct{
	Id uint32 `gorm:"primary_key"`
	Owner uint32 `json:"owner"`
	Url string `json:"url"`
	CreateTime string`json:"create_time";sql:"-"`
	Address string `json:"address"`
	Coordinate string `json:"coordinate"`
	Creator uint32 `json:"creator"`
	Sort uint32 `json:"sort"`
	Top uint8 `json:"top"`
	Score float64 `json:"score"`
	MachineComments []MachineComment `json:"machine_comments"`
}

