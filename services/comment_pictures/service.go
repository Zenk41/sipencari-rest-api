package comment_pictures

import (
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comment_pictures"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/comment_pictures"
	"github.com/Zenk41/sipencari-rest-api/models"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/comment_pictures"
)

type ComPictureService interface {
	Create(urlPictures []string, CommentID string) ([]response.CommentPicture, error)
	GetAll(CommentID string) ([]response.CommentPicture, error)
	GetByID(ComPictureID string) (response.CommentPicture, error)
	UpdateComPicture(payload payload.CommentPicture, comPictureID string) (response.CommentPicture, error)
	Delete(ComPictureID string) (bool, error)
}

type comPictureService struct {
	repository repository.ComPictureRepository
}

func NewComPictureService(repository repository.ComPictureRepository) ComPictureService {
	return &comPictureService{repository: repository}
}

func (cps *comPictureService) Create(urlPictures []string, CommentID string) ([]response.CommentPicture, error) {

	var rec []models.CommentPicture
	for _, pic := range urlPictures {
		rec = append(rec, models.CommentPicture{
			URL:       pic,
			CommentID: CommentID,
		})
	}
	pictures, err := cps.repository.Create(rec)
	if err != nil {
		return []response.CommentPicture{}, err
	}
	return *response.CommentPicturesResponse(pictures), nil
}
func (cps *comPictureService) GetAll(CommentID string) ([]response.CommentPicture, error) {
	comments, err := cps.repository.GetByCommentID(CommentID)

	if err != nil {
		return []response.CommentPicture{}, err
	}
	return *response.CommentPicturesResponse(comments), nil
}
func (cps *comPictureService) GetByID(ComPictureID string) (response.CommentPicture, error) {
	comPicture, err := cps.repository.GetByID(ComPictureID)

	if err != nil {
		return response.CommentPicture{}, err
	}
	return *response.CommentPictureResponse(comPicture), nil
}

func (cps *comPictureService) UpdateComPicture(payload payload.CommentPicture, comPictureID string) (response.CommentPicture, error) {
	picture, err := cps.repository.GetByID(comPictureID)
	if err != nil {
		return response.CommentPicture{}, err
	}
	picture.URL = payload.URL

	updatedPicture, err := cps.repository.Update(picture)

	if err != nil {
		return response.CommentPicture{}, err
	}
	return *response.CommentPictureResponse(updatedPicture), nil
}
func (cps *comPictureService) Delete(ComPictureID string) (bool, error) {
	isDeleted, err := cps.repository.Delete(ComPictureID)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}
