package auth

import (
	"GoToKube/config"
	"GoToKube/database"
	"GoToKube/logger"
	"GoToKube/web/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InitRootUser() error {
	db, _ := database.GetDBConnection()

	// 自动迁移模型
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	if config.Data.Auth.Pass == "" {
		err := errors.New("Root password is not set in environment variables")
		logger.GlobalLogger.Error(err.Error())
		return err
	}
	var user models.User
	if err := db.Where("username = ?", config.Data.Auth.User).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.Data.Auth.Pass), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user = models.User{
				Username: config.Data.Auth.User,
				Password: string(hashedPassword),
			}
			if err := db.Create(&user).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
