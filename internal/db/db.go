package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL 根据配置创建一个新的 GORM 实例
func NewMySQL(dsn string, maxIdle, maxOpen int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)

	return db, nil
}
