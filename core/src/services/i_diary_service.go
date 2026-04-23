// 日记服务接口定义
package services

import (
	"context"
	"go-diary-core/src/input_models"
	"go-diary-core/src/view_models"
)

// DiaryService 日记服务接口
// 定义日记的业务逻辑操作
type DiaryService interface {
	// CreateDiary 创建日记
	CreateDiary(ctx context.Context, input input_models.CreateDiaryInput) (*view_models.DiaryViewModel, error)

	// GetDiary 获取单条日记
	GetDiary(ctx context.Context, id int64) (*view_models.DiaryViewModel, error)

	// ListDiaries 获取日记列表
	ListDiaries(ctx context.Context) ([]*view_models.DiaryViewModel, error)

	// UpdateDiary 更新日记
	UpdateDiary(ctx context.Context, id int64, input input_models.UpdateDiaryInput) (*view_models.DiaryViewModel, error)

	// DeleteDiary 删除日记
	DeleteDiary(ctx context.Context, id int64) error
}
