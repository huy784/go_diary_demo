// 日记服务实现
package services

import (
	"context"
	"go-diary-core/src/input_models"
	"go-diary-core/src/models"
	"go-diary-core/src/view_models"
)

// diaryService 日记服务实现
type diaryService struct {
	repo DiaryRepository // 依赖日记仓库接口
}

// NewDiaryService 创建日记服务实例
func NewDiaryService(repo DiaryRepository) DiaryService {
	return &diaryService{repo: repo}
}

// CreateDiary 创建日记
func (s *diaryService) CreateDiary(ctx context.Context, input input_models.CreateDiaryInput) (*view_models.DiaryViewModel, error) {
	// 构建日记模型
	diary := &models.Diary{
		Title:   input.Title,
		Content: input.Content,
	}

	// 调用仓库层创建数据
	if err := s.repo.Create(ctx, diary); err != nil {
		return nil, err
	}

	// 转换为视图模型返回
	return s.toViewModel(diary), nil
}

// GetDiary 获取单条日记
func (s *diaryService) GetDiary(ctx context.Context, id int64) (*view_models.DiaryViewModel, error) {
	diary, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toViewModel(diary), nil
}

// ListDiaries 获取日记列表
func (s *diaryService) ListDiaries(ctx context.Context) ([]*view_models.DiaryViewModel, error) {
	diaries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为视图模型切片
	result := make([]*view_models.DiaryViewModel, len(diaries))
	for i, d := range diaries {
		result[i] = s.toViewModel(d)
	}
	return result, nil
}

// UpdateDiary 更新日记
func (s *diaryService) UpdateDiary(ctx context.Context, id int64, input input_models.UpdateDiaryInput) (*view_models.DiaryViewModel, error) {
	// 先获取原日记
	diary, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 更新字段（只更新非空字段）
	if input.Title != "" {
		diary.Title = input.Title
	}
	if input.Content != "" {
		diary.Content = input.Content
	}

	// 保存更新
	if err := s.repo.Update(ctx, diary); err != nil {
		return nil, err
	}

	return s.toViewModel(diary), nil
}

// DeleteDiary 删除日记
func (s *diaryService) DeleteDiary(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// toViewModel 将数据模型转换为视图模型
func (s *diaryService) toViewModel(d *models.Diary) *view_models.DiaryViewModel {
	return &view_models.DiaryViewModel{
		ID:        d.ID,
		Title:     d.Title,
		Content:   d.Content,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
