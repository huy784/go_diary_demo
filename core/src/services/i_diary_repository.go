// 日记仓库接口定义
package services

import (
	"context"
	"go-diary-core/src/models"
)

// DiaryRepository 日记仓库接口
// 定义日记数据的CRUD操作
type DiaryRepository interface {
	// Create 创建日记
	Create(ctx context.Context, diary *models.Diary) error

	// GetByID 根据ID获取日记
	GetByID(ctx context.Context, id string, userIdentity string) (*models.Diary, error)

	// List 获取所有日记
	List(ctx context.Context, userIdentity string) ([]*models.Diary, error)

	// Update 更新日记
	Update(ctx context.Context, diary *models.Diary) error

	// Delete 删除日记
	Delete(ctx context.Context, id string, userIdentity string) error
}
