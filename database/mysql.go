package database

import (
	"VDController/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	dbAddr string
	dbUser string
	dbPass string
	dbName string
}

// NewMySQLDB 创建一个新的MySQL数据库连接器
func NewMySQLDB(dbAddr, dbUser, dbPass, dbName string) *MySQLDB {
	return &MySQLDB{
		dbAddr: dbAddr,
		dbUser: dbUser,
		dbPass: dbPass,
		dbName: dbName,
	}
}

// Open 连接到MySQL数据库
func (m *MySQLDB) Open() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.dbUser, m.dbPass, m.dbAddr, m.dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
		return nil, err
	}
	return db, err
}

// Close 关闭MySQL数据库连接
func (m *MySQLDB) Close() error {
	return nil
}
