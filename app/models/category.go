package models

import "time"

type Category struct {
	ID           *string `json:"id" gorm:"type:char(36);primaryKey"`
	NamaCategory *string `json:"nama_category"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
