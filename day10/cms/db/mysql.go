package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	. "temp/model"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:xjv123..@tcp(127.0.0.1:3306)/go102?charset=utf8mb4&parseTime=True&loc=Local"

	// 创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 日志配置
	})

	// 迁移表
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})

}
