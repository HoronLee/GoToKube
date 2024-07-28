package database

import (
	"GoToKube/config"
	"GoToKube/logger"
	"gorm.io/gorm"
)

type Connector interface {
	Open() (*gorm.DB, error)
	Close() error
}

func CheckAndSetDefaultConfig() {
	// 如果 DBType 未设置，默认使用 sqlite
	if config.Data.DBType == "" {
		config.Data.DBType = "sqlite"
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
	}
	// 对于 SQLite，如果 DBPath 未设置，则设置为当前目录下的 data.db
	if config.Data.DBType == "sqlite" && config.Data.DBPath == "" {
		config.Data.DBPath = "./data.db"
		logger.GlobalLogger.Warn("SQLite database path is not set, defaulting to ./data.db")
	}
}

func CheckStatus() bool {
	CheckAndSetDefaultConfig()
	switch config.Data.DBType {
	case "":
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
		fallthrough
	case "sqlite":
		if config.Data.DBPath == "" {
			logger.GlobalLogger.Error("SQLite database path is not set")
			return false
		}
	case "mysql":
		if config.Data.DBAddr == "" || config.Data.DBUser == "" || config.Data.DBPass == "" || config.Data.DBName == "" {
			logger.GlobalLogger.Error("Missing required MySQL configuration parameters")
			return false
		}
	default:
		logger.GlobalLogger.Error("Unsupported database type: " + config.Data.DBType)
		return false
	}
	return true
}

func GetDBConnection() (db *gorm.DB, err error) {
	var dbHandler Connector
	switch config.Data.DBType {
	case "sqlite":
		dbHandler = NewSQLiteDB(config.Data.DBPath)
	case "mysql":
		dbHandler = NewMySQLDB(config.Data.DBAddr, config.Data.DBUser, config.Data.DBPass, config.Data.DBName)
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
