// 中间件包
package middleware

import (
	"go-diary-core/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserIdentityMiddleware 从 URL 路径提取用户标识
func UserIdentityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 URL 路径提取用户标识
		userIdentity := c.Param("userIdentity")

		// 验证用户标识是否在白名单中
		if userIdentity != "admin1" && userIdentity != "admin2" {
			response.Error(c, http.StatusForbidden, "Invalid user identity")
			c.Abort()
			return
		}

		// 将用户标识存储到上下文
		c.Set("userIdentity", userIdentity)
		c.Next()
	}
}
