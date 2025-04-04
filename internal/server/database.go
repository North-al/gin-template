package server

import (
	"fmt"
	"time"

	"github.com/North-al/gin-template/config"
	"github.com/North-al/gin-template/internal/data/models"
	"github.com/North-al/gin-template/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dbConfig := config.GetConfig().Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_", // 设置表名前缀
		},
	})

	if gin.Mode() == gin.DebugMode {
		// 在调试模式中自动迁移数据库
		db.AutoMigrate()
	}

	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get sql.DB: %v", err))
	}

	// 设置空闲连接池中最大连接数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping database: %v", err))
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %v", err))
	}

	DB = db

	logger.Info("数据库连接成功", "dsn：", dsn)
}
