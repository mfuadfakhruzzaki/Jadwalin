// middleware/auth.go
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
	"github.com/mfuadfakhruzaki/Jadwalin/utils"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Check if the token is blacklisted
        _, err := config.RedisClient.Get(context.Background(), "blacklist:"+tokenString).Result()
        if err != redis.Nil {
            // If the token is blacklisted
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
            c.Abort()
            return
        }

        // Validate JWT
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Set user_id in context
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}


