// 处理器层
// 处理HTTP请求，参数校验，调用服务层，返回响应
package handlers

import (
	"strconv"

	"go-diary-core/pkg/response"
	"go-diary-core/src/input_models"
	"go-diary-core/src/services"

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
// GET /api/v1/diaries
func (h *DiaryHandler) ListDiaries(c *gin.Context) {
	diaries, err := h.service.ListDiaries(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, diaries)
}

// GetDiary 获取单条日记
// GET /api/v1/diaries/:id
func (h *DiaryHandler) GetDiary(c *gin.Context) {
	// 解析路径参数id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	// 调用服务层获取日记
	diary, err := h.service.GetDiary(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "diary not found")
		return
	}
	response.Success(c, diary)
}

// CreateDiary 创建日记
// POST /api/v1/diaries
func (h *DiaryHandler) CreateDiary(c *gin.Context) {
	// 绑定并校验请求参数
	var input input_models.CreateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 调用服务层创建日记
	diary, err := h.service.CreateDiary(c.Request.Context(), input)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, diary)
}

// UpdateDiary 更新日记
// PUT /api/v1/diaries/:id
func (h *DiaryHandler) UpdateDiary(c *gin.Context) {
	// 解析路径参数id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	// 绑定并校验请求参数
	var input input_models.UpdateDiaryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 调用服务层更新日记
	diary, err := h.service.UpdateDiary(c.Request.Context(), id, input)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, diary)
}

// DeleteDiary 删除日记
// DELETE /api/v1/diaries/:id
func (h *DiaryHandler) DeleteDiary(c *gin.Context) {
	// 解析路径参数id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	// 调用服务层删除日记
	if err := h.service.DeleteDiary(c.Request.Context(), id); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, nil)
}
