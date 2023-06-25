package discussions

import (
	"errors"

	"github.com/Zenk41/sipencari-rest-api/constant"
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type discussionRepository struct {
	conn *gorm.DB
}

type DiscussionRepository interface {
	Create(Discussion models.Discussion, UrlPictures []string) (models.Discussion, error)
	GetAll(Page int, Size int, SortBy, Search, SearchQ, Privacy, Status string) (*gorm.DB, []models.Discussion, error)
	GetByID(DiscussionID string) (models.Discussion, error)
	GetByUserID(UserID string) ([]models.Discussion, error)
	GetAllWithoutStatus(Page int, Size int, SortBy, Search, SearchQ, Privacy string) (*gorm.DB, []models.Discussion, error)
	GetByUserIDWithPrivacy(UserID string, Privacy string) ([]models.Discussion, error)
	Update(Discussion models.Discussion) (models.Discussion, error)
	Delete(DiscussionID string) (bool, error)
}

func NewDiscussionRepository(conn *gorm.DB) DiscussionRepository {
	return &discussionRepository{
		conn: conn,
	}
}

func (dr *discussionRepository) Create(Discussion models.Discussion, UrlPictures []string) (models.Discussion, error) {
	Discussion.SetId(dr.conn)
	var disPictures []models.DiscussionPicture

	for _, pic := range UrlPictures {
		disPictures = append(disPictures, models.DiscussionPicture{
			URL:          pic,
			DiscussionID: Discussion.DiscussionID,
		})
	}

	Discussion.DiscussionPictures = disPictures

	result := dr.conn.
		Preload("User").
		Preload("DiscussionPictures").
		Preload("DiscussionLocation").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Comments.CommentLikes").
		Preload("Comments.CommentLikes.User").
		Preload("Comments.CommentReactions").
		Preload("Comments.CommentReactions.User").
		Preload("DiscussionLikes").
		Preload("DiscussionLikes.User").
		Create(&Discussion)
	result.Last(&Discussion)
	err := result.Error

	return Discussion, err
}

func (dr *discussionRepository) GetAll(Page int, Size int, SortBy, Search, SearchQ, Privacy, Status string) (*gorm.DB, []models.Discussion, error) {
	var rec []models.Discussion
	var model *gorm.DB

	if SearchQ != "" {
		model = dr.conn.Model(&rec).Order(SortBy).Where("privacy = ?", Privacy).Where("status = ?", Status).Where(SearchQ, Search)
		dr.conn.Offset(Page).Limit(Size).Order(SortBy).Where("privacy = ?", Privacy).Where("status = ?", Status).Where(SearchQ, Search).Preload("User").
			Preload("DiscussionPictures").
			Preload("DiscussionLocation").
			Preload("Comments").
			Preload("Comments.User").
			Preload("Comments.CommentLikes").
			Preload("Comments.CommentLikes.User").
			Preload("Comments.CommentReactions").
			Preload("Comments.CommentReactions.User").
			Preload("DiscussionLikes").
			Preload("DiscussionLikes.User").Find(&rec)
	} else {
		model = dr.conn.Model(&rec).Order(SortBy).Where("privacy = ?", Privacy).Where("status = ?", Status)
		dr.conn.Order(SortBy).Offset(Page).Limit(Size).Where("privacy = ?", Privacy).Where("status = ?", Status).Preload("User").
			Preload("DiscussionPictures").
			Preload("DiscussionLocation").
			Preload("Comments").
			Preload("Comments.User").
			Preload("Comments.CommentLikes").
			Preload("Comments.CommentLikes.User").
			Preload("Comments.CommentReactions").
			Preload("Comments.CommentReactions.User").
			Preload("DiscussionLikes").
			Preload("DiscussionLikes.User").Find(&rec)
	}

	return model, rec, nil
}

func (dr *discussionRepository) GetAllWithoutStatus(Page int, Size int, SortBy, Search, SearchQ, Privacy string) (*gorm.DB, []models.Discussion, error) {
	var rec []models.Discussion
	var model *gorm.DB

	if SearchQ != "" {
		model = dr.conn.Model(&rec).Order(SortBy).Where("privacy = ?", Privacy).Where(SearchQ, Search)
		dr.conn.Offset(Page).Limit(Size).Order(SortBy).Where("privacy = ?", Privacy).Where(SearchQ, Search).Preload("User").
			Preload("DiscussionPictures").
			Preload("DiscussionLocation").
			Preload("Comments").
			Preload("Comments.User").
			Preload("Comments.CommentLikes").
			Preload("Comments.CommentLikes.User").
			Preload("Comments.CommentReactions").
			Preload("Comments.CommentReactions.User").
			Preload("DiscussionLikes").
			Preload("DiscussionLikes.User").Find(&rec)
	} else {
		model = dr.conn.Model(&rec).Order(SortBy).Where("privacy = ?", Privacy)
		dr.conn.Order(SortBy).Offset(Page).Limit(Size).Where("privacy = ?", Privacy).Preload("User").
			Preload("DiscussionPictures").
			Preload("DiscussionLocation").
			Preload("Comments").
			Preload("Comments.User").
			Preload("Comments.CommentLikes").
			Preload("Comments.CommentLikes.User").
			Preload("Comments.CommentReactions").
			Preload("Comments.CommentReactions.User").
			Preload("DiscussionLikes").
			Preload("DiscussionLikes.User").Find(&rec)
	}

	return model, rec, nil
}

func (dr *discussionRepository) GetByID(DiscussionID string) (models.Discussion, error) {
	var rec models.Discussion
	error := dr.conn.
		Preload("User").
		Preload("DiscussionPictures").
		Preload("DiscussionLocation").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Comments.CommentLikes").
		Preload("Comments.CommentLikes.User").
		Preload("Comments.CommentReactions").
		Preload("Comments.CommentReactions.User").
		Preload("DiscussionLikes").
		Preload("DiscussionLikes.User").
		Where("discussions.discussion_id = ?", DiscussionID).First(&rec).Error
	return rec, error
}

func (dr *discussionRepository) GetByUserID(UserID string) ([]models.Discussion, error) {
	var rec []models.Discussion
	error := dr.conn.
		Preload("User").
		Preload("DiscussionPictures").
		Preload("DiscussionLocation").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Comments.CommentLikes").
		Preload("Comments.CommentLikes.User").
		Preload("Comments.CommentReactions").
		Preload("Comments.CommentReactions.User").
		Preload("DiscussionLikes").
		Preload("DiscussionLikes.User").Order("created_at desc").
		Where("user_id = ?", UserID).Find(&rec).Error
	return rec, error
}

func (dr *discussionRepository) GetByUserIDWithPrivacy(UserID string, Privacy string) ([]models.Discussion, error) {
	var rec []models.Discussion
	error := dr.conn.
		Preload("User").
		Preload("DiscussionPictures").
		Preload("DiscussionLocation").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Comments.CommentLikes").
		Preload("Comments.CommentLikes.User").
		Preload("Comments.CommentReactions").
		Preload("Comments.CommentReactions.User").
		Preload("DiscussionLikes").
		Preload("DiscussionLikes.User").
		Where("user_id = ?", UserID).Where("privacy = ?", Privacy).Find(&rec).Error
	return rec, error
}

func (dr *discussionRepository) Update(Discussion models.Discussion) (models.Discussion, error) {
	err := dr.conn.
		Preload("User").
		Preload("DiscussionPictures").
		Preload("DiscussionLocation").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Comments.CommentLikes").
		Preload("Comments.CommentLikes.User").
		Preload("Comments.CommentReactions").
		Preload("Comments.CommentReactions.User").
		Preload("DiscussionLikes").
		Preload("DiscussionLikes.User").
		Where("discussion_id = ?", Discussion.DiscussionID).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Updates(&Discussion).Error
	return Discussion, err
}

func (dr *discussionRepository) Delete(DiscussionID string) (bool, error) {
	rec, err := dr.GetByID(DiscussionID)
	if err != nil {
		return false, err
	}
	if rec.User.Role == constant.RoleSuperadmin {
		return false, errors.New("Forbidden")
	}
	if result := dr.conn.Delete(&rec); result.RowsAffected == 0 {
		return false, err
	}
	return true, nil
}
