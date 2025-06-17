package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 定义需要认证的路由标记
const RequireAuthKey = "require_auth"

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查当前路由是否需要认证
		if requireAuth, exists := c.Get(RequireAuthKey); exists && requireAuth.(bool) {
			// 执行认证逻辑
			token := c.GetHeader("Authorization")
			if token == "" || !validateToken(token) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// 需要认证的中间件 - 用于标记路由
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(RequireAuthKey, true)
		c.Next()
	}
}

// 模拟token验证
func validateToken(token string) bool {
	// 这里实现你的token验证逻辑
	return strings.HasPrefix(token, "Bearer valid-token")
}

// 业务处理函数
func publicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a public endpoint",
	})
}

func protectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a protected endpoint",
	})
}

func main() {
	r := gin.Default()

	// 全局使用认证中间件
	r.Use(AuthMiddleware())

	// 公开接口 - 不需要认证
	r.GET("/public", publicHandler)

	// 受保护的接口 - 需要认证
	r.GET("/protected", RequireAuth(), protectedHandler)

	// 也可以为路由组添加认证
	authGroup := r.Group("/api/v1")
	authGroup.Use(RequireAuth())
	{
		authGroup.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "User info"})
		})
		authGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "User profile"})
		})
	}

	r.Run(":8080")
}
