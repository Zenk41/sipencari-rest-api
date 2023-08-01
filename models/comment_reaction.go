package models

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"gorm.io/gorm"
)

type CommentReaction struct {
	ReactionID uint                 `json:"reaction_id"  gorm:"primaryKey"`
	UserID     string               `json:"user_id"`
	User       User                 `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Helpful    constant.HelpfulEnum `json:"helpful" gorm:"size:3;"`
	CommentID  string               `json:"comment_id"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	DeletedAt  gorm.DeletedAt       `json:"deleted_at" gorm:"index"`
}
