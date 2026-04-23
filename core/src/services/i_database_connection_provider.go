// 数据库连接提供者接口
package services

import "gorm.io/gorm"

// DatabaseConnectionProvider 数据库连接提供者接口
// 用于获取数据库连接实例
type DatabaseConnectionProvider interface {
	// GetDB 获取数据库连接实例
	GetDB() *gorm.DB
}
