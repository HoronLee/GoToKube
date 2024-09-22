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

	var user models.User
	// 尝试从数据库中查找 root 用户
	if err := db.Where("username = ?", config.Data.Auth.User).First(&user).Error; err != nil {
		// 如果 root 用户不存在，检查环境变量
		if err == gorm.ErrRecordNotFound {
			if config.Data.Auth.Pass == "" {
				// 如果没有设置环境变量中的密码，则返回错误
				err := errors.New("Root password is not set in environment variables")
				logger.GlobalLogger.Error(err.Error())
				return err
			}

			// 生成 bcrypt 密码哈希
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.Data.Auth.Pass), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			// 创建 root 用户
			user = models.User{
				Username: config.Data.Auth.User,
				Password: string(hashedPassword),
			}
			if err := db.Create(&user).Error; err != nil {
				return err
			}
		} else {
			// 处理其他数据库查询错误
			return err
		}
	} else {
		// 如果 root 用户已存在，记录日志并跳过
		logger.GlobalLogger.Info("Root user already exists, skipping creation.")
	}

	return nil
}
