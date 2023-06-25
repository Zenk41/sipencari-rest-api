package comments

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type commentRepository struct {
	conn *gorm.DB
}

type CommentRepository interface {
	Create(comment models.Comment, UrlPictures []string) (models.Comment, error)
	GetAll(discussionID string) ([]models.Comment, error)
	GetByID(commentID string) (models.Comment, error)
	Update(comment models.Comment, commentID string) (models.Comment, error)
	Delete(commentID string) (bool, error)
}

func NewCommentRepository(conn *gorm.DB) CommentRepository {
	return &commentRepository{
		conn: conn,
	}
}

func (cr *commentRepository) Create(Comment models.Comment, UrlPictures []string) (models.Comment, error) {
	Comment.SetId(cr.conn)
	var comPictures []models.CommentPicture
	for _, url := range UrlPictures {
		comPictures = append(comPictures, models.CommentPicture{
			URL:       url,
			CommentID: Comment.CommentID,
		})
	}
	Comment.CommentPictures = comPictures
	result := cr.conn.
		Preload("User").
		Preload("CommentLikes").
		Preload("CommentLikes.User").
		Preload("CommentPictures").
		Preload("CommentReactions").
		Preload("CommentReactions.User").
		Create(&Comment)
	result.Last(&Comment)
	err := result.Error

	return Comment, err
}

func (cr *commentRepository) GetAll(discussionID string) ([]models.Comment, error) {
	var rec []models.Comment
	err := cr.conn.
		Preload("User").
		Preload("CommentLikes").
		Preload("CommentLikes.User").
		Preload("CommentPictures").
		Preload("CommentReactions").
		Preload("CommentReactions.User").
		Where("discussion_id=?", discussionID).
		Find(&rec).Error
	return rec, err
}
func (cr *commentRepository) GetByID(commentID string) (models.Comment, error) {
	var rec models.Comment
	err := cr.conn.
		Preload("User").
		Preload("CommentLikes").
		Preload("CommentLikes.User").
		Preload("CommentReactions").
		Preload("CommentPictures").
		Preload("CommentReactions.User").
		Where("comment_id=?", commentID).
		First(&rec).Error
	return rec, err
}
func (cr *commentRepository) Update(comment models.Comment, commentID string) (models.Comment, error) {
	err := cr.conn.
	Session(&gorm.Session{FullSaveAssociations: true}).
		Preload("User").
		Preload("CommentLikes").
		Preload("CommentLikes.User").
		Preload("CommentPictures").
		Preload("CommentReactions").
		Preload("CommentReactions.User").
		Model(&comment).
		Where("comment_id = ?", comment.CommentID).
		Select("message").Updates(&comment).Error

	return comment, err
}

func (cr *commentRepository) Delete(commentID string) (bool, error) {
	rec, err := cr.GetByID(commentID)
	if err != nil {
		return false, err
	}
	if result := cr.conn.Unscoped().Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
