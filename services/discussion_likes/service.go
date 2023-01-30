package discussion_likes

import (
	response "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_likes"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/discussion_likes"
)

type DisLikeService interface {
	GetByID(discussionID string, userID string) (response.DiscussionLike, error)
	GetAll(discussionID string) ([]response.DiscussionLike, error)
	Like(discussionID string, userID string) (bool, string, error)
}

type disLikeService struct {
	repository repository.DisLikeRepository
}

func NewDisLikeService(repository repository.DisLikeRepository) DisLikeService {
	return &disLikeService{repository: repository}
}

func (dls *disLikeService) GetByID(discussionID string, userID string) (response.DiscussionLike, error) {
	like, err := dls.repository.GetLike(discussionID, userID)
	if err != nil {
		return response.DiscussionLike{}, err
	}
	return *response.DiscussionLikeResponse(like), err
}

func (dls *disLikeService) GetAll(discussionID string) ([]response.DiscussionLike, error) {
	likes, err := dls.repository.GetAll(discussionID)
	if err != nil {
		return []response.DiscussionLike{}, err
	}
	return *response.DiscussionLikesResponse(likes), err
}

func (dls *disLikeService) Like(discussionID string, userID string) (bool, string, error) {
	var isDeleted bool
	var err error
	isLike, err := dls.repository.GetLike(discussionID, userID)
	if isLike.UserID != userID {
		_, err = dls.repository.Like(discussionID, userID)
		if err != nil {
			return false, "", err
		}
		return true, "like", nil
	}
	err = nil
	isDeleted, err = dls.repository.DeleteLike(discussionID, userID)
	if err != nil {
		return isDeleted, "", err
	}
	return isDeleted, "unlike", nil

}
