package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/valeelim/mahchat/pkg/repository"
)

func Authorized(cache repository.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		values := strings.Split(authHeader, " ")
		if len(values) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization header",
			})
		}
		bearer, token := values[0], values[1]
		if bearer != "Basic" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid bearer type",
			})
		}
		if _, err := cache.GetAccessToken(context.Background(), token); err == redis.Nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
		}
		c.Next()
	}
}
