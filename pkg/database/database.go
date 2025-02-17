package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"schaub-dev.xyz/fabrice/fiber-template/pkg"
)

type DBClient struct {
	conn *gorm.DB
}

func NewDbClient(config *pkg.Config, _ *fiber.App) (*DBClient, error) {
	var db *gorm.DB
	var err error
	attempts := 0
	for attempts < config.DbConnRetries {
		db, err = gorm.Open(mysql.Open(config.DatabaseUrl), &gorm.Config{})
		if err == nil {
			break
		}
		attempts++
		timeToWait := time.Duration(attempts) * time.Second
		time.Sleep(timeToWait)
	}

	if err != nil {
		return nil, err
	}

	//register all repositories here
	db.AutoMigrate()

	return &DBClient{conn: db}, nil
}
