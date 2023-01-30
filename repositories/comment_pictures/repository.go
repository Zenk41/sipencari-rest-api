package comment_pictures

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type comPictureRepository struct {
	conn *gorm.DB
}

type ComPictureRepository interface {
	Create(ComPictures []models.CommentPicture) ([]models.CommentPicture, error)
	GetByID(ComPictureID string) (models.CommentPicture, error)
	GetByCommentID(CommentID string) ([]models.CommentPicture, error)
	Update(ComPicture models.CommentPicture) (models.CommentPicture, error)
	Delete(ComPictureID string) (bool, error)
}

func NewComPictureRepository(conn *gorm.DB) ComPictureRepository {
	return &comPictureRepository{
		conn: conn,
	}
}

func (cpr *comPictureRepository) Create(ComPictures []models.CommentPicture) ([]models.CommentPicture, error) {
	err := cpr.conn.Create(&ComPictures).Error
	return ComPictures, err
}

func (cpr *comPictureRepository) GetByID(ComPictureID string) (models.CommentPicture, error) {
	var rec models.CommentPicture
	error := cpr.conn.Where("picture_id = ?", ComPictureID).First(&rec).Error
	return rec, error
}

func (cpr *comPictureRepository) GetByCommentID(CommentID string) ([]models.CommentPicture, error) {
	var rec []models.CommentPicture
	error := cpr.conn.Where("comment_id = ?", CommentID).Find(&rec).Error
	return rec, error
}

func (cpr *comPictureRepository) Update(ComPicture models.CommentPicture) (models.CommentPicture, error) {
	err := cpr.conn.Save(&ComPicture).Error
	return ComPicture, err
}
func (cpr *comPictureRepository) Delete(ComPictureID string) (bool, error) {
	rec, err := cpr.GetByID(ComPictureID)
	if err != nil {
		return false, err
	}
	if result := cpr.conn.Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
