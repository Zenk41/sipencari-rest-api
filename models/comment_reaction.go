package models

import (
	"time"

	"gorm.io/gorm"
)

type CommentReaction struct {
	UserID     string         `json:"user_id" gorm:"primaryKey"`
	User       User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Helpful    bool           `json:"helpful"`
	NotHelp    bool           `json:"nothelp"`
	CommentID  uint           `json:"comment_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
