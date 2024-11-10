package database

import (
	"GoToKube/config"
	"GoToKube/logger"
	"errors"

	"gorm.io/gorm"
)

type Connector interface {
	Open() (*gorm.DB, error)
	Close() error
}

func CheckAndSetDefaultConfig() {
	// 如果 Type 未设置，默认使用 sqlite
	if config.Data.Type == "" {
		config.Data.Type = "sqlite"
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
	}
	// 对于 SQLite，如果Path 未设置，则设置为当前目录下的 data.db
	if config.Data.Type == "sqlite" && config.Data.Path == "" {
		config.Data.Path = "./data.db"
		logger.GlobalLogger.Warn("SQLite database path is not set, defaulting to ./data.db")
	}
}

func CheckStatus() error {
	CheckAndSetDefaultConfig()
	switch config.Data.Type {
	case "":
		logger.GlobalLogger.Warn("Database type is not specified, defaulting to sqlite")
		fallthrough
	case "sqlite":
		if config.Data.Path == "" {
			return errors.New("SQLite database path is not set")
		}
	case "mysql":
		if config.Data.Addr == "" || config.Data.Database.User == "" || config.Data.Password == "" || config.Data.Name == "" {
			logger.GlobalLogger.Error("Missing required MySQL configuration parameters")
			return errors.New("Missing required MySQL configuration parameters")
		}
	default:
		return errors.New("Unsupported database type: " + config.Data.Type)
	}
	return nil
}

func GetDBConnection() (db *gorm.DB, err error) {
	var dbHandler Connector
	switch config.Data.Type {
	case "sqlite":
		dbHandler = NewSQLiteDB(config.Data.Path)
	case "mysql":
		dbHandler = NewMySQLDB(config.Data.Addr, config.Data.Database.User, config.Data.Password, config.Data.Name)
	default:
		logger.GlobalLogger.Error("Unsupported database type")
		panic("Unsupported database type")
	}
	db, err = dbHandler.Open()
	if err != nil {
		logger.GlobalLogger.Error("Unable to Connect to database")
		return nil, err
	}
	return db, err
}
