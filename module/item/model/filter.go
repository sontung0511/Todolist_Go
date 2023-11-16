package model

type Filter struct {
	Status string `json:"status" gorm:"column:status;"`
}
