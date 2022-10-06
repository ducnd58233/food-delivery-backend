package common

import "time"

type SQLModel struct {
	Id        int `json:"id" gorm:"column:id;"`
	Status    int `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"created_at" gorm:"created_at"`
}