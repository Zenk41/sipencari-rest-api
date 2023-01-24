package discussions

import "gorm.io/gorm"

type discussionRepository struct {
	conn *gorm.DB
}

type DiscussionRepository interface {

}

func NewDiscussionRepository(conn *gorm.DB) DiscussionRepository {
	return &discussionRepository{
		conn: conn,
	}
}