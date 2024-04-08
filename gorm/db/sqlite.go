package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBSqlite(f string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(f), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbdb, _ := db.DB()
	dbdb.SetMaxIdleConns(512)
	dbdb.SetMaxOpenConns(512)
	dbdb.SetConnMaxLifetime(3 * time.Minute)

	return db, nil
}

func CreateTable(db *gorm.DB, table ...interface{}) error {
	return db.AutoMigrate(table...)
}
