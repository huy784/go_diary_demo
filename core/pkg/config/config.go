// 配置管理模块
package config

import "os"

// Config 应用配置结构体
// 用于存储应用程序的所有配置项
type Config struct {
	Port     string // HTTP服务端口
	DBHost   string // 数据库主机地址
	DBPort   string // 数据库端口
	DBUser   string // 数据库用户名
	DBPass   string // 数据库密码
	DBName   string // 数据库名称
}

// Load 从环境变量加载配置
// 如果环境变量未设置，则使用默认值
// 返回加载后的配置指针
func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   getEnv("DB_PORT", "5432"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPass:   getEnv("DB_PASS", "postgres"),
		DBName:   getEnv("DB_NAME", "diarydb"),
	}
}

// getEnv 获取环境变量值
// 如果环境变量存在则返回其值，否则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
