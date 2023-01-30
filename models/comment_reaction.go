package models

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"gorm.io/gorm"
)

type CommentReaction struct {
	UserID    string               `json:"user_id" gorm:"primaryKey"`
	User      User                 `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Helpful   constant.HelpfulEnum `json:"helpful"`
	CommentID string                 `json:"comment_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `json:"deleted_at" gorm:"index"`
}
