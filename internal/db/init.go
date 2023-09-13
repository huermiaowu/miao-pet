package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pet/internal/config"
)

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		return err
	}

	// 根据结构自动建表
	err = DB.AutoMigrate(&Pet{})
	if err != nil {
		return err
	}
	return nil
}
