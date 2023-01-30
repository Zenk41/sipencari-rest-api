package feedbacks

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type feedbackRepository struct {
	conn *gorm.DB
}

type FeedbackRepository interface {
	Create(Feedback models.Feedback) (models.Feedback, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []models.Feedback, error)
	GetByID(FeedbackID string) (models.Feedback, error)
	GetByUserID(UserID string) ([]models.Feedback, error)
	GetByEmail(UserEmail string) ([]models.Feedback, error)
	Update(Feedback models.Feedback) (models.Feedback, error)
	Delete(UserID string) (bool, error)
}

func NewFeedbackRepository(conn *gorm.DB) FeedbackRepository {
	return &feedbackRepository{
		conn: conn,
	}
}

func (fr *feedbackRepository) Create(Feedback models.Feedback) (models.Feedback, error) {
	Feedback.SetId(fr.conn)
	result := fr.conn.Preload("User").Create(&Feedback).Last(&Feedback)
	err := result.Error
	return Feedback, err
}
func (fr *feedbackRepository) GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []models.Feedback, error) {
	var rec []models.Feedback
	var model *gorm.DB
	if SortBy == "" && SearchQ == "" {
		model = fr.conn.Model(&rec)
	} else if SortBy != "" && SearchQ == "" {
		model = fr.conn.Order(SortBy).Model(&rec)
	} else if SortBy == "" && SearchQ != "" {
		model = fr.conn.Model(&rec).Where(SearchQ, Search)
	} else {
		model = fr.conn.Order(SortBy).Model(&rec).Where(SearchQ, Search)
	}

	if err := model.Offset(Page).Preload("User").Limit(Size).Find(&rec).Error; err != nil {
		return model, []models.Feedback{}, err
	}

	return model, rec, nil
}

func (fr *feedbackRepository) GetByID(FeedbackID string) (models.Feedback, error) {
	var rec models.Feedback
	error := fr.conn.Preload("User").Where("feedback_id = ?", FeedbackID).First(&rec).Error
	return rec, error
}

func (fr *feedbackRepository) GetByUserID(UserID string) ([]models.Feedback, error) {
	var rec []models.Feedback
	error := fr.conn.Preload("User").Where("User.UserID = ?", UserID).Find(&rec).Error
	return rec, error
}

func (fr *feedbackRepository) GetByEmail(UserEmail string) ([]models.Feedback, error) {
	var rec []models.Feedback
	error := fr.conn.Preload("User").Where("User.Email = ?", UserEmail).Find(&rec).Error
	return rec, error
}

func (fr *feedbackRepository) Update(Feedback models.Feedback) (models.Feedback, error) {
	err := fr.conn.Preload("User").Save(&Feedback).Error
	return Feedback, err
}

func (fr *feedbackRepository) Delete(FeedbackID string) (bool, error) {
	rec, err := fr.GetByID(FeedbackID)
	if err != nil {
		return false, err
	}
	if result := fr.conn.Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
