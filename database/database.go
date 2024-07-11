package database

import (
	"GoToKube/config"
	"GoToKube/logger"
	"gorm.io/gorm"
)

type DatabaseConnector interface {
	Open() (*gorm.DB, error)
	Close() error
}

var (
	dbconfig = config.ConfigData
)

func CheckAndSetDefaultConfig() {
	// 如果 DBType 未设置，默认使用 sqlite
	if dbconfig.DBType == "" {
		dbconfig.DBType = "sqlite"
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
	}
	// 对于 SQLite，如果 DBPath 未设置，则设置为当前目录下的 data.db
	if dbconfig.DBType == "sqlite" && dbconfig.DBPath == "" {
		dbconfig.DBPath = "./data.db"
		logger.GlobalLogger.Warn("SQLite database path is not set, defaulting to ./data.db")
	}
}

func CheckStatus() bool {
	CheckAndSetDefaultConfig()
	switch dbconfig.DBType {
	case "":
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
		fallthrough
	case "sqlite":
		if dbconfig.DBPath == "" {
			logger.GlobalLogger.Error("SQLite database path is not set")
			return false
		}
	case "mysql":
		if dbconfig.DBAddr == "" || dbconfig.DBUser == "" || dbconfig.DBPass == "" || dbconfig.DBName == "" {
			logger.GlobalLogger.Error("Missing required MySQL configuration parameters")
			return false
		}
	default:
		logger.GlobalLogger.Error("Unsupported database type: " + dbconfig.DBType)
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
