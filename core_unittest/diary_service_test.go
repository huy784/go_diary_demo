package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go-diary-core/src/input_models"
	"go-diary-core/src/models"
	"go-diary-core/src/services"
)

// MockDiaryRepository 是 DiaryRepository 接口的 mock 实现
type MockDiaryRepository struct {
	diaries map[string]*models.Diary
}

// NewMockDiaryRepository 创建一个新的 MockDiaryRepository
func NewMockDiaryRepository() *MockDiaryRepository {
	return &MockDiaryRepository{
		diaries: make(map[string]*models.Diary),
	}
}

// Create 实现 DiaryRepository 接口的 Create 方法
func (m *MockDiaryRepository) Create(ctx context.Context, diary *models.Diary) error {
	// 生成 UUID
	if diary.ID == "" {
		diary.ID = uuid.New().String()
	}
	m.diaries[diary.ID] = diary
	return nil
}

// GetByID 实现 DiaryRepository 接口的 GetByID 方法
func (m *MockDiaryRepository) GetByID(ctx context.Context, id string, userIdentity string) (*models.Diary, error) {
	diary, exists := m.diaries[id]
	if !exists || diary.UserIdentityGuid != userIdentity {
		return nil, nil
	}
	return diary, nil
}

// List 实现 DiaryRepository 接口的 List 方法
func (m *MockDiaryRepository) List(ctx context.Context, userIdentity string) ([]*models.Diary, error) {
	var diaries []*models.Diary
	for _, diary := range m.diaries {
		if diary.UserIdentityGuid == userIdentity {
			diaries = append(diaries, diary)
		}
	}
	return diaries, nil
}

// Update 实现 DiaryRepository 接口的 Update 方法
func (m *MockDiaryRepository) Update(ctx context.Context, diary *models.Diary) error {
	m.diaries[diary.ID] = diary
	return nil
}

// Delete 实现 DiaryRepository 接口的 Delete 方法
func (m *MockDiaryRepository) Delete(ctx context.Context, id string, userIdentity string) error {
	if diary, exists := m.diaries[id]; exists && diary.UserIdentityGuid == userIdentity {
		delete(m.diaries, id)
	}
	return nil
}

// TestDiaryService_CreateDiary 测试创建日记功能
func TestDiaryService_CreateDiary(t *testing.T) {
	// 准备测试数据
	repo := NewMockDiaryRepository()
	service := services.NewDiaryService(repo)
	ctx := context.Background()
	userIdentity := "admin1"

	input := input_models.CreateDiaryInput{
		Title:   "Test Diary",
		Content: "This is a test diary",
	}

	// 执行测试
	result, err := service.CreateDiary(ctx, userIdentity, input)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, input.Title, result.Title)
	assert.Equal(t, input.Content, result.Content)
}

// TestDiaryService_GetDiary 测试获取日记功能
func TestDiaryService_GetDiary(t *testing.T) {
	// 准备测试数据
	repo := NewMockDiaryRepository()
	service := services.NewDiaryService(repo)
	ctx := context.Background()
	userIdentity := "admin1"

	// 先创建一个日记
	input := input_models.CreateDiaryInput{
		Title:   "Test Diary",
		Content: "This is a test diary",
	}
	createdDiary, _ := service.CreateDiary(ctx, userIdentity, input)

	// 执行测试
	result, err := service.GetDiary(ctx, userIdentity, createdDiary.ID)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdDiary.ID, result.ID)
	assert.Equal(t, createdDiary.Title, result.Title)
	assert.Equal(t, createdDiary.Content, result.Content)
}

// TestDiaryService_ListDiaries 测试获取日记列表功能
func TestDiaryService_ListDiaries(t *testing.T) {
	// 准备测试数据
	repo := NewMockDiaryRepository()
	service := services.NewDiaryService(repo)
	ctx := context.Background()
	userIdentity := "admin1"

	// 创建两个日记
	input1 := input_models.CreateDiaryInput{
		Title:   "Test Diary 1",
		Content: "This is test diary 1",
	}
	input2 := input_models.CreateDiaryInput{
		Title:   "Test Diary 2",
		Content: "This is test diary 2",
	}
	// 确保两个日记都被创建
	_, err1 := service.CreateDiary(ctx, userIdentity, input1)
	_, err2 := service.CreateDiary(ctx, userIdentity, input2)
	assert.NoError(t, err1)
	assert.NoError(t, err2)

	// 执行测试
	result, err := service.ListDiaries(ctx, userIdentity)

	// 验证结果
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

// TestDiaryService_UpdateDiary 测试更新日记功能
func TestDiaryService_UpdateDiary(t *testing.T) {
	// 准备测试数据
	repo := NewMockDiaryRepository()
	service := services.NewDiaryService(repo)
	ctx := context.Background()
	userIdentity := "admin1"

	// 先创建一个日记
	input := input_models.CreateDiaryInput{
		Title:   "Test Diary",
		Content: "This is a test diary",
	}
	createdDiary, _ := service.CreateDiary(ctx, userIdentity, input)

	// 准备更新数据
	updateInput := input_models.UpdateDiaryInput{
		Title:   "Updated Test Diary",
		Content: "This is an updated test diary",
	}

	// 执行测试
	result, err := service.UpdateDiary(ctx, userIdentity, createdDiary.ID, updateInput)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdDiary.ID, result.ID)
	assert.Equal(t, updateInput.Title, result.Title)
	assert.Equal(t, updateInput.Content, result.Content)
}

// TestDiaryService_DeleteDiary 测试删除日记功能
func TestDiaryService_DeleteDiary(t *testing.T) {
	// 准备测试数据
	repo := NewMockDiaryRepository()
	service := services.NewDiaryService(repo)
	ctx := context.Background()
	userIdentity := "admin1"

	// 先创建一个日记
	input := input_models.CreateDiaryInput{
		Title:   "Test Diary",
		Content: "This is a test diary",
	}
	createdDiary, _ := service.CreateDiary(ctx, userIdentity, input)

	// 执行测试
	err := service.DeleteDiary(ctx, userIdentity, createdDiary.ID)

	// 验证结果
	assert.NoError(t, err)

	// 验证日记已被删除
	deletedDiary, err := service.GetDiary(ctx, userIdentity, createdDiary.ID)
	assert.NoError(t, err)
	assert.Nil(t, deletedDiary)
}
