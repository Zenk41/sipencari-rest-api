package discussion_pictures

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type disPictureRepository struct {
	conn *gorm.DB
}

type DisPictureRepository interface {
	Create(DisPictures []models.DiscussionPicture) ([]models.DiscussionPicture, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []models.DiscussionPicture, error)
	GetByID(DisPictureID string) (models.DiscussionPicture, error)
	GetByDiscussionID(DiscussionID string) ([]models.DiscussionPicture, error)
	Update(DisPicture models.DiscussionPicture) (models.DiscussionPicture, error)
	Delete(DisPictureID string) (bool, error)
}

func NewDisPictureRepository(conn *gorm.DB) DisPictureRepository {
	return &disPictureRepository{
		conn: conn,
	}
}
func (dpr *disPictureRepository) Create(DisPictures []models.DiscussionPicture) ([]models.DiscussionPicture, error) {
	err := dpr.conn.Create(&DisPictures).Error
	return DisPictures, err
}

func (dpr *disPictureRepository) GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []models.DiscussionPicture, error) {
	var rec []models.DiscussionPicture
	var model *gorm.DB
	if SortBy == "" && SearchQ == "" {
		model = dpr.conn.Model(&rec)
	} else if SortBy != "" && SearchQ == "" {
		model = dpr.conn.Model(&rec).Order(SortBy)
	} else {
		model = dpr.conn.Model(&rec).Order(SortBy).Where(SearchQ, Search)
	}

	if err := model.Find(&rec); err != nil {
		return model, []models.DiscussionPicture{}, err.Error
	}

	return model, rec, nil
}

func (dpr *disPictureRepository) GetByID(DisPictureID string) (models.DiscussionPicture, error) {
	var rec models.DiscussionPicture
	error := dpr.conn.Where("picture_id = ?", DisPictureID).First(&rec).Error
	return rec, error
}

func (dpr *disPictureRepository) GetByDiscussionID(DiscussionID string) ([]models.DiscussionPicture, error) {
	var rec []models.DiscussionPicture
	error := dpr.conn.Where("discussion_id = ?", DiscussionID).Find(&rec).Error
	return rec, error
}

func (dpr *disPictureRepository) Update(DisPicture models.DiscussionPicture) (models.DiscussionPicture, error) {
	err := dpr.conn.Updates(DisPicture).Error
	return DisPicture, err
}

func (dpr *disPictureRepository) Delete(DisPictureID string) (bool, error) {
	rec, err := dpr.GetByID(DisPictureID)
	if err != nil {
		return false, err
	}
	if result := dpr.conn.Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
