package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscussionLocation struct {
	LocationID   uint           `json:"location_id" gorm:"primaryKey"`
	Lat          float64        `json:"lat"`
	Lng          float64        `json:"lng"`
	LocationName string         `json:"location_name"`
	DiscussionID string         `json:"discussion_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
