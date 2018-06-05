package model

type Machine struct{
	Id uint32 `gorm:"primary_key"`
	OwnerId uint32 `json:"owner_id"`
	LbsId string `json:"lbs_id"`
	CreateTime string`json:"create_time" sql:"-"`
	Address string `json:"address"`
	CreatorId uint32 `json:"creator_id"`
	Lat string `json:"Lat"`
	Lon string `json:"lon"`
	Cover string `json:"cover"`
	No string `json:"no"`
	AuthId string `json:"auth_id"`
	Sort uint32 `json:"sort"`
	Top uint8 `json:"top"`
	Score float64 `json:"score"`
	MachineComments []MachineComment `json:"machine_comments"`
}

