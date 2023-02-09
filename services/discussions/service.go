package discussions

import (
	"strings"

	"github.com/Zenk41/sipencari-rest-api/constant"
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussions"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/discussions"
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/discussions"
)

type DiscussionService interface {
	Create(payload payload.CreateDiscussion, UserID string, UrlPictures []string, locationName string) (response.Discussion, error)
	GetAll(Page int, Size int, SortBy, Status, Privacy, Search, SearchQ string) (*gorm.DB, []response.Discussion, error)
	GetByID(discussionID string) (response.Discussion, error)
	GetByUserID(userID string, privacy string) ([]response.Discussion, error)
	GetMyDiscussion(userID string) ([]response.Discussion, error)
	Update(payload payload.UpdateDiscussion, discussionID string, locationName string) (response.Discussion, error)
	Delete(discussionID string) (bool, error)
}

type discussionService struct {
	repository repository.DiscussionRepository
}

func NewDiscussionService(repository repository.DiscussionRepository) DiscussionService {
	return &discussionService{repository: repository}
}

func (ds *discussionService) Create(payload payload.CreateDiscussion, UserID string, UrlPictures []string, locationName string) (response.Discussion, error) {

	Discussion := models.Discussion{
		UserID:   UserID,
		Title:    payload.Title,
		Category: constant.CategoryEnum(payload.Category),
		Content:  payload.Content,
		DiscussionLocation: models.DiscussionLocation{
			Lat:          payload.Lat,
			Lng:          payload.Lng,
			LocationName: locationName,
		},
		Status:  constant.StatusEnum(payload.Status),
		Privacy: constant.PrivacyEnum(payload.Privacy),
	}

	discussion, err := ds.repository.Create(Discussion, UrlPictures)
	if err != nil {
		return response.Discussion{}, err
	}

	return *response.DiscussionResponse(discussion), err
}

func (ds *discussionService) GetMyDiscussion(userID string) ([]response.Discussion, error) {
	var discussions []models.Discussion
	var err error

	discussions, err = ds.repository.GetByUserID(userID)
	if err != nil {
		return []response.Discussion{}, err
	}
	return *response.DiscussionsResponse(discussions), nil
}

func (ds *discussionService) GetAll(Page int, Size int, SortBy, Status, Privacy, Search, SearchQ string) (*gorm.DB, []response.Discussion, error) {
	var sort string
	var search string
	var searchQ string
	var model *gorm.DB
	var discussions []models.Discussion
	var err error

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
	if Status != "" {
		model, discussions, err = ds.repository.GetAll(Page, Size, sort, search, searchQ, Privacy, Status)
	} else {
		model, discussions, err = ds.repository.GetAllWithoutStatus(Page, Size, sort, search, searchQ, Privacy)
	}

	if err != nil {
		return model, []response.Discussion{}, err
	}

	return model, *response.DiscussionsResponse(discussions), nil
}

func (ds *discussionService) GetByID(discussionID string) (response.Discussion, error) {
	discussion, err := ds.repository.GetByID(discussionID)
	if err != nil {
		return response.Discussion{}, err
	}
	return *response.DiscussionResponse(discussion), nil
}

func (ds *discussionService) GetByUserID(userID string, privacy string) ([]response.Discussion, error) {
	var discussions []models.Discussion
	var err error

	if privacy == "" {
		discussions, err = ds.repository.GetByUserID(userID)
	} else {
		discussions, err = ds.repository.GetByUserIDWithPrivacy(userID, privacy)
	}
	if err != nil {
		return []response.Discussion{}, err
	}
	return *response.DiscussionsResponse(discussions), nil
}

func (ds *discussionService) Update(payload payload.UpdateDiscussion, discussionID string, locationName string) (response.Discussion, error) {
	discussion, err := ds.repository.GetByID(discussionID)
	if err != nil {
		return response.Discussion{}, err
	}
	discussion.Title = payload.Title
	discussion.Category = constant.CategoryEnum(payload.Category)
	discussion.Content = payload.Content
	discussion.DiscussionLocation.Lat = payload.Lat
	discussion.DiscussionLocation.Lng = payload.Lng
	discussion.Status = constant.StatusEnum(payload.Status)
	discussion.Privacy = constant.PrivacyEnum(payload.Privacy)
	discussion.DiscussionLocation.LocationName = locationName

	updatedDiscussion, err := ds.repository.Update(discussion)
	if err != nil {
		return response.Discussion{}, err
	}
	return *response.DiscussionResponse(updatedDiscussion), nil
}

func (ds *discussionService) Delete(DiscussionID string) (bool, error) {
	isDeleted, err := ds.repository.Delete(DiscussionID)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}
