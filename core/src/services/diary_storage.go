// 日记仓库实现
package services

import (
	"context"
	"go-diary-core/src/models"

	"gorm.io/gorm"
)

// diaryRepository 日记仓库实现
type diaryRepository struct {
	db *gorm.DB
}

// NewDiaryRepository 创建日记仓库实例
func NewDiaryRepository(dbProvider DatabaseConnectionProvider) DiaryRepository {
	return &diaryRepository{db: dbProvider.GetDB()}
}

// Create 创建日记记录
func (r *diaryRepository) Create(ctx context.Context, diary *models.Diary) error {
	return r.db.WithContext(ctx).Create(diary).Error
}

// GetByID 根据ID获取日记
func (r *diaryRepository) GetByID(ctx context.Context, id string, userIdentity string) (*models.Diary, error) {
	var diary models.Diary
	err := r.db.WithContext(ctx).Where("id = ? AND user_identity_guid = ?", id, userIdentity).First(&diary).Error
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

// List 获取所有日记
func (r *diaryRepository) List(ctx context.Context, userIdentity string) ([]*models.Diary, error) {
	var diaries []*models.Diary
	err := r.db.WithContext(ctx).Where("user_identity_guid = ?", userIdentity).Find(&diaries).Error
	return diaries, err
}

// Update 更新日记
func (r *diaryRepository) Update(ctx context.Context, diary *models.Diary) error {
	return r.db.WithContext(ctx).Save(diary).Error
}

// Delete 删除日记
func (r *diaryRepository) Delete(ctx context.Context, id string, userIdentity string) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_identity_guid = ?", id, userIdentity).Delete(&models.Diary{}).Error
}
