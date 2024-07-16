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

var (
	dbConfig = config.Data
)

func CheckAndSetDefaultConfig() {
	// 如果 DBType 未设置，默认使用 sqlite
	if dbConfig.DBType == "" {
		dbConfig.DBType = "sqlite"
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
	}
	// 对于 SQLite，如果 DBPath 未设置，则设置为当前目录下的 data.db
	if dbConfig.DBType == "sqlite" && dbConfig.DBPath == "" {
		dbConfig.DBPath = "./data.db"
		logger.GlobalLogger.Warn("SQLite database path is not set, defaulting to ./data.db")
	}
}

func CheckStatus() bool {
	CheckAndSetDefaultConfig()
	switch dbConfig.DBType {
	case "":
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
		fallthrough
	case "sqlite":
		if dbConfig.DBPath == "" {
			logger.GlobalLogger.Error("SQLite database path is not set")
			return false
		}
	case "mysql":
		if dbConfig.DBAddr == "" || dbConfig.DBUser == "" || dbConfig.DBPass == "" || dbConfig.DBName == "" {
			logger.GlobalLogger.Error("Missing required MySQL configuration parameters")
			return false
		}
	default:
		logger.GlobalLogger.Error("Unsupported database type: " + dbConfig.DBType)
		return false
	}
	return true
}

func GetDBConnection() (db *gorm.DB, err error) {
	var dbHandler Connector
	switch dbConfig.DBType {
	case "sqlite":
		dbHandler = NewSQLiteDB(dbConfig.DBPath)
	case "mysql":
		dbHandler = NewMySQLDB(dbConfig.DBAddr, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName)
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
