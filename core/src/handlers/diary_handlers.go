// 处理器层
// 处理HTTP请求，参数校验，调用服务层，返回响应
package handlers

import (
	"go-diary-core/pkg/response"
	"go-diary-core/src/input_models"
	"go-diary-core/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DiaryHandler 日记处理器
type DiaryHandler struct {
	service services.DiaryService // 依赖服务层接口
}

// NewDiaryHandler 创建日记处理器实例
func NewDiaryHandler(service services.DiaryService) *DiaryHandler {
	return &DiaryHandler{service: service}
}

// ListDiaries 获取日记列表
func (h *DiaryHandler) ListDiaries(c *gin.Context) {
	// 从上下文获取用户标识
	userIdentity, _ := c.Get("userIdentity")

	// 调用服务层获取日记列表
	diaries, err := h.service.ListDiaries(c.Request.Context(), userIdentity.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get diaries")
		return
	}

	// 返回成功响应
	response.Success(c, diaries)
}

// GetDiary 获取单条日记
func (h *DiaryHandler) GetDiary(c *gin.Context) {
	// 从上下文获取用户标识
	userIdentity, _ := c.Get("userIdentity")

	// 获取路径参数
	id := c.Param("id")

	// 调用服务层获取日记
	diary, err := h.service.GetDiary(c.Request.Context(), userIdentity.(string), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Diary not found")
		return
	}

	// 返回成功响应
	response.Success(c, diary)
}

// CreateDiary 创建日记
func (h *DiaryHandler) CreateDiary(c *gin.Context) {
	// 从上下文获取用户标识
	userIdentity, _ := c.Get("userIdentity")

	// 绑定请求参数
	var input input_models.CreateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// 调用服务层创建日记
	diary, err := h.service.CreateDiary(c.Request.Context(), userIdentity.(string), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create diary")
		return
	}

	// 返回成功响应
	response.Success(c, diary)
}

// UpdateDiary 更新日记
func (h *DiaryHandler) UpdateDiary(c *gin.Context) {
	// 从上下文获取用户标识
	userIdentity, _ := c.Get("userIdentity")

	// 获取路径参数
	id := c.Param("id")

	// 绑定请求参数
	var input input_models.UpdateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// 调用服务层更新日记
	diary, err := h.service.UpdateDiary(c.Request.Context(), userIdentity.(string), id, input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update diary")
		return
	}

	// 返回成功响应
	response.Success(c, diary)
}

// DeleteDiary 删除日记
func (h *DiaryHandler) DeleteDiary(c *gin.Context) {
	// 从上下文获取用户标识
	userIdentity, _ := c.Get("userIdentity")

	// 获取路径参数
	id := c.Param("id")

	// 调用服务层删除日记
	err := h.service.DeleteDiary(c.Request.Context(), userIdentity.(string), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete diary")
		return
	}

	// 返回成功响应
	response.Success(c, nil)
}
