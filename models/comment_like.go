package models

import (
	"time"

	"gorm.io/gorm"
)

type CommentLike struct {
	UserID    string         `json:"user_id"`
	User      User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentID string         `json:"comment_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
