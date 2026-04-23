// 统一响应格式模块
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
// 所有API响应都使用此格式
type Response struct {
	Code    int         `json:"code"`    // 状态码，0表示成功
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data,omitempty"` // 响应数据
}

// Success 返回成功响应
// c: Gin上下文
// data: 响应数据
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 返回错误响应
// c: Gin上下文
// code: 错误码
// message: 错误消息
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

// BadRequest 返回400错误（请求参数错误）
func BadRequest(c *gin.Context, message string) {
	Error(c, 400, message)
}

// NotFound 返回404错误（资源不存在）
func NotFound(c *gin.Context, message string) {
	Error(c, 404, message)
}

// InternalServerError 返回500错误（服务器内部错误）
func InternalServerError(c *gin.Context, message string) {
	Error(c, 500, message)
}
