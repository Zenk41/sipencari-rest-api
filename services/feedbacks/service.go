package feedbacks

import (
	"strings"

	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/feedbacks"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/feedbacks"
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/feedbacks"
)

type FeedbackService interface {
	Create(payload payload.Feedback, UserID string) (response.Feedback, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []response.Feedback, error)
	GetByID(feedbackID string) (response.Feedback, error)
	GetByUserID(userID string) ([]response.Feedback, error)
	GetByEmail(userEmail string) ([]response.Feedback, error)
	UpdateFeedback(payload payload.Feedback, feedbackID string) (response.Feedback, error)
	Delete(feedbackID string) (bool, error)
}

type feedbackService struct {
	repository repository.FeedbackRepository
}

func NewFeedbackService(repository repository.FeedbackRepository) FeedbackService {
	return &feedbackService{repository: repository}
}

func (fs *feedbackService) Create(payload payload.Feedback, UserID string) (response.Feedback, error) {
	Feedback := models.Feedback{
		UserID:   UserID,
		Reaction: payload.Reaction,
		Review:   payload.Review,
	}

	feedback, err := fs.repository.Create(Feedback)
	if err != nil {
		return response.Feedback{}, err
	}

	return *response.FeedbackResponse(feedback), nil
}

func (fs *feedbackService) GetAll(Page int, Size int, SortBy, Search, SearchQ string) (*gorm.DB, []response.Feedback, error) {
	var sort string
	var search string
	var searchQ string


	if SearchQ == "" {
		searchQ = ""
	} else {
		searchQ = SearchQ + " Like ? "
		search = "%" + Search + "%"
	}

	if SortBy != "" {
		if strings.HasPrefix(SortBy, "-") {
			sort = SortBy[1:] + " DESC"
		} else {
			sort = SortBy[0:] + " ASC"
		}
	} else {
		sort = ""
	}
	model, feedbacks, err := fs.repository.GetAll(Page, Size, sort, search, searchQ)
	if err != nil {
		return model, []response.Feedback{}, err
	}

	return model, *response.FeedbacksResponse(feedbacks), nil
}
func (fs *feedbackService) GetByID(feedbackID string) (response.Feedback, error) {
	feedback, err := fs.repository.GetByID(feedbackID)
	if err != nil {
		return response.Feedback{}, err
	}
	return *response.FeedbackResponse(feedback), nil
}
func (fs *feedbackService) GetByUserID(userID string) ([]response.Feedback, error) {
	feedbacks, err := fs.repository.GetByUserID(userID)
	if err != nil {
		return []response.Feedback{}, err
	}
	return *response.FeedbacksResponse(feedbacks), nil
}
func (fs *feedbackService) GetByEmail(userEmail string) ([]response.Feedback, error) {
	feedbacks, err := fs.repository.GetByEmail(userEmail)
	if err != nil {
		return []response.Feedback{}, err
	}
	return *response.FeedbacksResponse(feedbacks), nil
}
func (fs *feedbackService) UpdateFeedback(payload payload.Feedback, feedbackID string) (response.Feedback, error) {
	feedback, err := fs.repository.GetByID(feedbackID)
	if err != nil {
		return response.Feedback{}, err
	}
	feedback.Reaction = payload.Reaction
	feedback.Review = payload.Review

	updatedFeedback, err := fs.repository.Update(feedback)
	if err != nil {
		return response.Feedback{}, err
	}
	return *response.FeedbackResponse(updatedFeedback), nil
}
func (fs *feedbackService) Delete(feedbackID string) (bool, error) {
	isDeleted, err := fs.repository.Delete(feedbackID)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}
