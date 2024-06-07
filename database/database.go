package database

import (
	"VDController/config"
	"VDController/logger"
	"VDController/web/models"
	"errors"
	"gorm.io/gorm"
)

type DatabaseConnector interface {
	Open() (*gorm.DB, error)
	Close() error
}

var (
	dbconfig = config.ConfigData
)

func CheckStatus() bool {
	requiredParams := map[string]bool{
		"sqlite":  dbconfig.DBPath != "",
		"mysql":   dbconfig.DBAddr != "" && dbconfig.DBUser != "" && dbconfig.DBPass != "" && dbconfig.DBName != "",
		"default": false,
	}
	paramValid := requiredParams[dbconfig.DBType]
	if !paramValid {
		logger.GlobalLogger.Error("Missing required database configuration parameters")
		return false
	}
	return true
}

func GetDBConnection() (db *gorm.DB, err error) {
	var dbHandler DatabaseConnector
	switch dbconfig.DBType {
	case "sqlite":
		dbHandler = NewSQLiteDB(dbconfig.DBPath)
	case "mysql":
		dbHandler = NewMySQLDB(dbconfig.DBAddr, dbconfig.DBUser, dbconfig.DBPass, dbconfig.DBName)
	default:
		logger.GlobalLogger.Error("Unsupported database type")
		panic("Unsupported database type")
	}
	db, err = dbHandler.Open()
	if err != nil {
		logger.GlobalLogger.Error("Unable to Connect to database")
		panic(err)
	}
	return db, err
}

func SaveOrUpdateStatusInfo(info models.StatusInfo) error {
	db, err := GetDBConnection()
	db.AutoMigrate(&models.StatusInfo{})
	if err != nil {
		return err
	}
	// 查询是否存在相同记录
	var existingInfo models.StatusInfo
	err = db.Where("component = ?", info.Component).First(&existingInfo).Error
	if err == nil { // 记录已存在，进行更新
		existingInfo.Version = info.Version
		existingInfo.Status = info.Status
		err = db.Save(&existingInfo).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) { // 记录不存在，创建新记录
		err = db.Create(&info).Error
	} else {
		return err
	}
	return err
}
