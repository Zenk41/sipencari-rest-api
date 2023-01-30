package discussion_locations

import (
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
)

type disLocationRepository struct {
	conn *gorm.DB
}

type DisLocationRepository interface {
	GetAll() ([]models.DiscussionLocation, error)
	GetByID(DisLocationID string) (models.DiscussionLocation, error)
	GetByDiscussionID(DiscussionID string) (models.DiscussionLocation, error)
	Update(DisLocation models.DiscussionLocation) (models.DiscussionLocation, error)
}

func NewDisLocationRepository(conn *gorm.DB) DisLocationRepository {
	return &disLocationRepository{
		conn: conn,
	}
}

func (dlr *disLocationRepository) GetAll() ([]models.DiscussionLocation, error) {
	var rec []models.DiscussionLocation

	err := dlr.conn.Find(&rec).Error

	return rec, err
}

func (dlr *disLocationRepository) GetByID(DisLocationID string) (models.DiscussionLocation, error) {
	var rec models.DiscussionLocation

	err := dlr.conn.Where("location_id = ?", DisLocationID).First(&rec).Error

	return rec, err
}

func (dlr *disLocationRepository) GetByDiscussionID(DiscussionID string) (models.DiscussionLocation, error) {
	var rec models.DiscussionLocation

	err := dlr.conn.Where("discussion_id = ?" ,DiscussionID).First(&rec).Error

	return rec, err
}
func (dlr *disLocationRepository) Update(DisLocation models.DiscussionLocation) (models.DiscussionLocation, error) {

	err := dlr.conn.Save(&DisLocation).Error

	return DisLocation, err

}
