package entity

import "time"

type IDModels struct {
	ID uint `gorm:"autoIncrement:true"`
}

type TimeModels struct {
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
