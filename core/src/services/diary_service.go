// 日记服务实现
package services

import (
	"context"
	"go-diary-core/src/input_models"
	"go-diary-core/src/models"
	"go-diary-core/src/view_models"

	"gorm.io/gorm"
)

// diaryService 日记服务实现
type diaryService struct {
	repo DiaryRepository
}

// NewDiaryService 创建日记服务实例
func NewDiaryService(repo DiaryRepository) DiaryService {
	return &diaryService{repo: repo}
}

// CreateDiary 创建日记
func (s *diaryService) CreateDiary(ctx context.Context, userIdentity string, input input_models.CreateDiaryInput) (*view_models.DiaryViewModel, error) {
	// 创建日记模型
	diary := &models.Diary{
		UserIdentityGuid: userIdentity,
		Title:            input.Title,
		Content:          input.Content,
	}

	// 调用仓库创建
	err := s.repo.Create(ctx, diary)
	if err != nil {
		return nil, err
	}

	// 转换为视图模型
	return &view_models.DiaryViewModel{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
	}, nil
}

// GetDiary 根据ID获取日记
func (s *diaryService) GetDiary(ctx context.Context, userIdentity string, id string) (*view_models.DiaryViewModel, error) {
	// 调用仓库获取
	diary, err := s.repo.GetByID(ctx, id, userIdentity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// 检查日记是否存在
	if diary == nil {
		return nil, nil
	}

	// 转换为视图模型
	return &view_models.DiaryViewModel{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
	}, nil
}

// ListDiaries 获取所有日记
func (s *diaryService) ListDiaries(ctx context.Context, userIdentity string) ([]*view_models.DiaryViewModel, error) {
	// 调用仓库获取
	diaries, err := s.repo.List(ctx, userIdentity)
	if err != nil {
		return nil, err
	}

	// 转换为视图模型
	viewModels := make([]*view_models.DiaryViewModel, len(diaries))
	for i, diary := range diaries {
		viewModels[i] = &view_models.DiaryViewModel{
			ID:        diary.ID,
			Title:     diary.Title,
			Content:   diary.Content,
			CreatedAt: diary.CreatedAt,
			UpdatedAt: diary.UpdatedAt,
		}
	}

	return viewModels, nil
}

// UpdateDiary 更新日记
func (s *diaryService) UpdateDiary(ctx context.Context, userIdentity string, id string, input input_models.UpdateDiaryInput) (*view_models.DiaryViewModel, error) {
	// 先获取日记
	diary, err := s.repo.GetByID(ctx, id, userIdentity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// 检查日记是否存在
	if diary == nil {
		return nil, nil
	}

	// 更新字段
	if input.Title != "" {
		diary.Title = input.Title
	}
	if input.Content != "" {
		diary.Content = input.Content
	}

	// 调用仓库更新
	err = s.repo.Update(ctx, diary)
	if err != nil {
		return nil, err
	}

	// 转换为视图模型
	return &view_models.DiaryViewModel{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
	}, nil
}

// DeleteDiary 删除日记
func (s *diaryService) DeleteDiary(ctx context.Context, userIdentity string, id string) error {
	return s.repo.Delete(ctx, id, userIdentity)
}