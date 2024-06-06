package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	dbPath string
}

func NewSQLiteDB(dbPath string) *SQLiteDB {
	return &SQLiteDB{dbPath: dbPath}
}

func (s *SQLiteDB) Open() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.dbPath), &gorm.Config{})
	return db, err
}

func (s *SQLiteDB) Close() error {
	return nil
}
