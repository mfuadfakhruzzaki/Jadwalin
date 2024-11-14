package config

import "github.com/gin-gonic/gin"

var (
    ServerPort = GetEnv("APP_PORT", "8080")
    CorsOrigins = GetEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
)

func InitServer() *gin.Engine {
    r := gin.Default()

    // Setup CORS middleware
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", CorsOrigins)
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Next()
    })

    return r
}
