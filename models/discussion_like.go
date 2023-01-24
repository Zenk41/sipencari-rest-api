package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscussionLike struct {
	UserID           string         `json:"user_id" gorm:"primaryKey"`
	User             User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DiscussionLikeID uint           `json:"discussion_like_id" `
	DiscussionID     string         `json:"discussion_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}