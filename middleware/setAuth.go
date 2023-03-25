package middleware

import (
	"chatgpt-go/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetAuthorizationHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := global.OpenAIKey // 使用从环境变量或输入中获取的API密钥
		c.Request.Header.Set("Authorization", "Bearer "+token)
		c.Next()
	}
}

func isNotEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authSecretKey := global.Config.System.AuthSecretKey
		if isNotEmptyString(authSecretKey) {
			authorization := c.GetHeader("Authorization")
			if authorization == "" || strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer ")) != strings.TrimSpace(authSecretKey) {
				response := struct {
					Status  string      `json:"status"`
					Message string      `json:"message"`
					Data    interface{} `json:"data"`
				}{
					Status:  "Unauthorized",
					Message: "Error: 无访问权限 | No access rights",
					Data:    nil,
				}
				c.JSON(http.StatusUnauthorized, response)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
