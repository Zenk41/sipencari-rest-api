package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	CommentID       string            `json:"comment_id" gorm:"size:255;primaryKey"`
	Message         string            `json:"message"`
	DiscussionID    string            `json:"discussion_id"`
	CommentPictures []CommentPicture  `json:"comment_pictures" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentLikes    []CommentLike     `json:"comment_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentReactions []CommentReaction `json:"comment_reactions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// ParrentComment  string            `json:"parrent_comment"`
	UserID          string            `json:"user_id"`
	User            User              `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

func (c *Comment) SetId(db *gorm.DB) {
	c.CommentID = uuid.New().String()
	return
}
