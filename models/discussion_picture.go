package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscussionPicture struct {
	PictureID    uint           `json:"picture_id" gorm:"primaryKey"`
	URL          string         `json:"url"`
	DiscussionID string         `json:"discussion_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
