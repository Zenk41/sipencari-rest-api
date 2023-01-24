package sql

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Zenk41/sipencari-rest-api/config"
	"github.com/Zenk41/sipencari-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB perform connecting to a Database Server
func InitDB() *gorm.DB {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	dbConfig := config.LoadDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.HOST,
		dbConfig.USERNAME,
		dbConfig.PASSWORD,
		dbConfig.NAME,
		dbConfig.PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("error when connecting to a database server: %s", err))
	}

	log.Print("connected to a database server")

	return db
}

func MigrationDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.CommentPicture{},
		&models.CommentLike{},
		&models.CommentReaction{},
		&models.DiscussionPicture{},
		&models.DiscussionLocation{},
		&models.DiscussionLike{},
		&models.Discussion{},
		&models.Comment{},
		&models.Feedback{},
	)
}

// CloseDB perform closing a Database Server
func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance : %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection : %v", err)
		return err
	}
	log.Print("database connection is closed")
	return nil
}
