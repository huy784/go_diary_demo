// 主程序入口
package main

import (
	"log"

	"go-diary-core/pkg/config"
	"go-diary-core/src/handlers"
	"go-diary-core/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库连接
	dbProvider := services.NewDatabaseConnectionProvider(cfg)

	// 初始化仓库层（数据访问层）
	diaryRepo := services.NewDiaryRepository(dbProvider)

	// 初始化服务层（业务逻辑层）
	diaryService := services.NewDiaryService(diaryRepo)

	// 初始化处理器层（HTTP处理层）
	diaryHandler := handlers.NewDiaryHandler(diaryService)

	// 创建Gin默认实例（包含Logger和Recovery中间件）
	r := gin.Default()

	// API路由分组
	api := r.Group("/api/v1")
	{
		// 日记相关路由
		diaries := api.Group("/diaries")
		{
			diaries.GET("", diaryHandler.ListDiaries)        // 获取日记列表
			diaries.GET("/:id", diaryHandler.GetDiary)      // 获取单条日记
			diaries.POST("", diaryHandler.CreateDiary)      // 创建日记
			diaries.PUT("/:id", diaryHandler.UpdateDiary)  // 更新日记
			diaries.DELETE("/:id", diaryHandler.DeleteDiary) // 删除日记
		}
	}

	// 启动HTTP服务器
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
