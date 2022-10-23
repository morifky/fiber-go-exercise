package database

import (
	"fiber-go-exercise/pkg/models"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		zap.S().Warn("Cannot connect to database: ", DbName)
	} else {
		zap.S().Info("We are connected to: ", DbName)
	}

	return postgresDB, err

}

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
}
