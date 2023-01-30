package comment_likes

import (
	response "github.com/Zenk41/sipencari-rest-api/dto/response/comment_likes"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/comment_likes"
)

type ComLikeService interface {
	GetByID(commentID string, userID string) (response.CommentLike, error)
	GetAll(commentID string) ([]response.CommentLike, error)
	Like(commentID string, userID string) (bool, string, error)
}

type comLikeService struct {
	repository repository.ComLikeRepository
}

func NewComLikeService(repository repository.ComLikeRepository) ComLikeService {
	return &comLikeService{repository: repository}
}

func (cls *comLikeService) GetByID(commentID string, userID string) (response.CommentLike, error) {
	like, err := cls.repository.GetLike(commentID, userID)
	if err != nil {
		return response.CommentLike{}, err
	}
	return *response.CommentLikeResponse(like), err
}
func (cls *comLikeService) GetAll(commentID string) ([]response.CommentLike, error) {
	likes, err := cls.repository.GetAll(commentID)
	if err != nil {
		return []response.CommentLike{}, err
	}
	return *response.CommentLikesResponse(likes), err
}
func (cls *comLikeService) Like(commentID string, userID string) (bool, string, error) {
	var isDeleted bool
	var err error
	isLike, err := cls.repository.GetLike(commentID, userID)
	if isLike.UserID != userID {
		_, err = cls.repository.Like(commentID, userID)
		if err != nil {
			return false, "", err
		}
		return true, "like", nil
	}
	err = nil
	isDeleted, err = cls.repository.DeleteLike(commentID, userID)
	if err != nil {
		return isDeleted, "", err
	}
	return isDeleted, "unlike", nil
}
