package models

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Discussion struct {
	DiscussionID       string                `json:"discussion_id" gorm:"size:255;primaryKey"`
	Title              string                `json:"title"`
	Category           constant.CategoryEnum `json:"category" gorm:"size:5"`
	Content            string                `json:"content"`
	DiscussionPictures []DiscussionPicture   `json:"discussion_pictures" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DiscussionLocation DiscussionLocation    `json:"discussion_location" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DiscussionLikes    []DiscussionLike      `json:"discussion_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID             string                `json:"user_id"`
	User               User                  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments           []Comment             `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status             constant.StatusEnum   `json:"status" gorm:"size:12"`
	Type               constant.TypeEnum     `json:"type" gorm:"size:5"`
	Privacy            constant.PrivacyEnum  `json:"privacy" gorm:"size:7"`
	Contact            string                `json:"contact" gorm:"size:13"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
	DeletedAt          gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
}

func (d *Discussion) SetId(db *gorm.DB) {
	d.DiscussionID = uuid.New().String()
	return
}
