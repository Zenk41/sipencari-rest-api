package comment_reactions

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type comReactionRepository struct {
	conn *gorm.DB
}

type ComReactionRepository interface {
	Create(reaction models.CommentReaction) (models.CommentReaction, error)
	Update(reaction models.CommentReaction, commentID string) (models.CommentReaction, error)
	Delete(commentID, userID string) (bool, error)
	GetByID(commentID, userID string) (models.CommentReaction, error)
	GetAll(commentID string) ([]models.CommentReaction, error)
}

func NewComReactionRepository(conn *gorm.DB) ComReactionRepository {
	return &comReactionRepository{
		conn: conn,
	}
}

func (crr *comReactionRepository) Create(reaction models.CommentReaction) (models.CommentReaction, error) {
	result := crr.conn.Preload("User").Create(&reaction)
	result.Last(&reaction)
	err := result.Error
	return reaction, err

}
func (crr *comReactionRepository) Update(reaction models.CommentReaction, commentID string) (models.CommentReaction, error) {
	err := crr.conn.Preload("User").Where("comment_id=?", commentID).Updates(&reaction).Error
	return reaction, err
}

func (crr *comReactionRepository) Delete(commentID, userID string) (bool, error) {
	rec, err := crr.GetByID(commentID, userID)
	if err != nil {
		return false, err
	}
	if result := crr.conn.Where("comment_id=? AND user_id=?", commentID, userID).Unscoped().Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}

func (crr *comReactionRepository) GetByID(commentID, userID string) (models.CommentReaction, error) {
	var rec models.CommentReaction
	error := crr.conn.Preload("User").Where("comment_id = ? AND user_id = ?", commentID, userID).First(&rec).Error
	return rec, error
}

func (crr *comReactionRepository) GetAll(commentID string) ([]models.CommentReaction, error) {
	var rec []models.CommentReaction
	err := crr.conn.Where("comment_id = ?", commentID).Preload("User").
		Find(&rec).Error

	return rec, err
}
