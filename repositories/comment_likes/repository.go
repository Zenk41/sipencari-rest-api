package comment_likes

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type comLikeRepository struct {
	conn *gorm.DB
}

type ComLikeRepository interface {
	GetAll(commentID string) ([]models.CommentLike, error)
	GetLike(commentID string, userID string) (models.CommentLike, error)
	Like(commentID string, userID string) (models.CommentLike, error)
	DeleteLike(commentID string, userID string) (bool, error)
}

func NewComLikeRepository(conn *gorm.DB) ComLikeRepository {
	return &comLikeRepository{
		conn: conn,
	}
}

func (clr *comLikeRepository) GetAll(commentID string) ([]models.CommentLike, error) {
	var rec []models.CommentLike
	err := clr.conn.Preload("User").Where("comment_id = ?", commentID).Find(&rec).Error
	return rec, err
}

func (clr *comLikeRepository) GetLike(commentID string, userID string) (models.CommentLike, error) {
	var rec models.CommentLike

	err := clr.conn.Preload("User").Where("comment_id=? AND user_id=?", commentID, userID).First(&rec).Error

	return rec, err
}

func (clr *comLikeRepository) Like(commentID string, userID string) (models.CommentLike, error) {
	var rec models.CommentLike

	rec.UserID = userID
	rec.CommentID = commentID

	result := clr.conn.Preload("User").Create(&rec).Last(&rec)
	err := result.Error

	return rec, err
}

func (clr *comLikeRepository) DeleteLike(commentID string, userID string) (bool, error) {
	rec, err := clr.GetLike(commentID, userID)
	if err != nil || rec.UserID == "" {
		return false, err
	}

	if result := clr.conn.Where("comment_id=? AND user_id=?", commentID, userID).Unscoped().Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
