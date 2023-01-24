package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	FeedbackID string         `json:"feedback_id" gorm:"primaryKey"`
	UserID     string         `json:"user_id"`
	User       User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reaction   string         `json:"reaction"`
	Review     string         `json:"review"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (f *Feedback) SetId(db *gorm.DB) {
	f.FeedbackID = uuid.New().String()
	return
}
