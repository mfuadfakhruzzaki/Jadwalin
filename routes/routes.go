// routes/routes.go
package routes

import (
	"github.com/mfuadfakhruzaki/Jadwalin/controllers"
	middleware "github.com/mfuadfakhruzaki/Jadwalin/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Routes for Authentication (Tidak perlukan middleware)
    r.POST("/auth/register", controllers.Register)
    r.POST("/auth/login", controllers.Login)
    r.POST("/auth/logout", controllers.Logout)
    r.POST("/auth/reset-password", controllers.ResetPassword)

    // Routes for FCM Token (Protected)
    protected := r.Group("/api")
    protected.Use(middleware.AuthRequired()) // Menggunakan middleware AuthRequired
    {
        protected.POST("/fcm/token", controllers.SaveFCMToken)
        protected.DELETE("/fcm/token/:token", controllers.DeleteFCMToken)
    }

    // Routes for Course (Protected)
    protected.POST("/courses", controllers.CreateCourse)
    protected.GET("/courses/:user_id", controllers.GetUserCourses)
}