package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"create_at" gorm:"column:create_at;"`
	UpdateAt  *time.Time `json:"update_at,omitempty" gorm:"column:update_at;"`
}
