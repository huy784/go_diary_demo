// 数据库连接提供者实现
package services

import (
	"fmt"
	"log"

	"go-diary-core/pkg/config"
	"go-diary-core/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// databaseConnectionProvider 数据库连接提供者实现
type databaseConnectionProvider struct {
	db *gorm.DB
}

// NewDatabaseConnectionProvider 创建数据库连接提供者
// 根据配置连接数据库并自动迁移表结构
func NewDatabaseConnectionProvider(cfg *config.Config) DatabaseConnectionProvider {
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	// 连接到数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表结构
	if err := db.AutoMigrate(&models.Diary{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return &databaseConnectionProvider{db: db}
}

// GetDB 获取数据库连接实例
func (p *databaseConnectionProvider) GetDB() *gorm.DB {
	return p.db
}
