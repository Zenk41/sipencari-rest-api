package discussion_pictures

import (
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussion_pictures"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_pictures"
	"github.com/Zenk41/sipencari-rest-api/models"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/discussion_pictures"
)

type DisPictureService interface {
	Create(urlPictures []string, DiscussionID string) ([]response.DiscussionPicture, error)
	GetAll(DiscussionID string) ([]response.DiscussionPicture, error)
	GetByID(DisPictureID string) (response.DiscussionPicture, error)
	UpdateDisPicture(payload payload.DiscussionPicture, disPictureID string) (response.DiscussionPicture, error)
	Delete(DisPictureID string) (bool, error)
}

type disPictureService struct {
	repository repository.DisPictureRepository
}

func NewDisPictureService(repository repository.DisPictureRepository) DisPictureService {
	return &disPictureService{repository: repository}
}

func (dps *disPictureService) Create(urlPictures []string, DiscussionID string) ([]response.DiscussionPicture, error) {
	var rec []models.DiscussionPicture
	for _, pic := range urlPictures {
		rec = append(rec, models.DiscussionPicture{
			URL:          pic,
			DiscussionID: DiscussionID,
		})
	}
	pictures, err := dps.repository.Create(rec)
	if err != nil {
		return []response.DiscussionPicture{}, err
	}
	return *response.DiscussionPicturesResponse(pictures), nil

}

func (dps *disPictureService) GetAll(DiscussionID string) ([]response.DiscussionPicture, error) {

	discussions, err := dps.repository.GetByDiscussionID(DiscussionID)

	if err != nil {
		return []response.DiscussionPicture{}, err
	}
	return *response.DiscussionPicturesResponse(discussions), nil
}

func (dps *disPictureService) GetByID(DisPictureID string) (response.DiscussionPicture, error) {
	disPicture, err := dps.repository.GetByID(DisPictureID)

	if err != nil {
		return response.DiscussionPicture{}, err
	}
	return *response.DiscussionPictureResponse(disPicture), nil
}

func (dps *disPictureService) UpdateDisPicture(payload payload.DiscussionPicture, disPictureID string) (response.DiscussionPicture, error) {
	picture, err := dps.repository.GetByID(disPictureID)
	if err != nil {
		return response.DiscussionPicture{}, err
	}
	picture.URL = payload.URL

	updatedPicture, err := dps.repository.Update(picture)

	if err != nil {
		return response.DiscussionPicture{}, err
	}
	return *response.DiscussionPictureResponse(updatedPicture), nil
}

func (dps *disPictureService) Delete(DisPictureID string) (bool, error) {
	isDeleted, err := dps.repository.Delete(DisPictureID)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}
