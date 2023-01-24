package models

import (
	"time"

	"gorm.io/gorm"
)

type CommentPicture struct {
	PictureID uint           `json:"picture_id" gorm:"primaryKey"`
	URL       string         `json:"url"`
	Comment   Comment        `json:"comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentID string         `json:"comment_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
