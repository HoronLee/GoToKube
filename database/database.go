package database

import (
	"VDController/config"
	"VDController/logger"
	"gorm.io/gorm"
)

type DatabaseConnector interface {
	Open() (*gorm.DB, error)
	Close() error
}

var globalDB *gorm.DB

func CheckStatus() {
	dbconfig := config.ConfigData
	// 检查是否所有必需的数据库连接参数都存在
	if dbconfig.DBType == "" || (dbconfig.DBType == "mysql" && (dbconfig.DBAddr == "" || dbconfig.DBUser == "" || dbconfig.DBPass == "" || dbconfig.DBName == "")) {
		logger.GlobalLogger.Error("Missing required database configuration parameters")
		panic("Missing required database configuration parameters")
	} else {
		InitDB()
	}
}

func InitDB() {
	dbconfig := config.ConfigData
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
	db, err := dbHandler.Open()
	if err != nil {
		panic(err)
	}
	globalDB = db
}
