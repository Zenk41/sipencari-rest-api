package discussion_likes

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type disLikeRepository struct {
	conn *gorm.DB
}

type DisLikeRepository interface {
	GetAll(discussionID string) ([]models.DiscussionLike, error)
	GetLike(discussionID string, userID string) (models.DiscussionLike, error)
	Like(discussionID string, userID string) (models.DiscussionLike, error)
	DeleteLike(discussionID string, userID string) (bool, error)
}

func NewDisLikeRepository(conn *gorm.DB) DisLikeRepository {
	return &disLikeRepository{
		conn: conn,
	}
}

func (dlr *disLikeRepository) GetAll(discussionID string) ([]models.DiscussionLike, error) {
	var rec []models.DiscussionLike
	err := dlr.conn.Where("discussion_id = ?", discussionID).Preload("User").
		Find(&rec).Error

	return rec, err

}

func (dlr *disLikeRepository) GetLike(discussionID string, userID string) (models.DiscussionLike, error) {
	var rec models.DiscussionLike

	err := dlr.conn.Preload("User").Where("discussion_id=? AND user_id=?", discussionID, userID).First(&rec).Error

	return rec, err

}

func (dlr *disLikeRepository) Like(discussionID string, userID string) (models.DiscussionLike, error) {
	var rec models.DiscussionLike

	rec.UserID = userID
	rec.DiscussionID = discussionID

	result := dlr.conn.Preload("User").Create(&rec).Last(&rec)
	err := result.Error

	return rec, err
}

func (dlr disLikeRepository) DeleteLike(discussionID string, userID string) (bool, error) {
	rec, err := dlr.GetLike(discussionID, userID)
	if err != nil || rec.UserID == "" {
		return false, err
	}

	if result := dlr.conn.Where("discussion_id=? AND user_id=?", discussionID, userID).Unscoped().Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
