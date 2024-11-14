package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
	"github.com/mfuadfakhruzaki/Jadwalin/models"
)

// SaveFCMToken akan menyimpan atau memperbarui token FCM untuk pengguna yang sudah terautentikasi
func SaveFCMToken(c *gin.Context) {
    var input struct {
        Token string `json:"token" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
        return
    }

    // Get the user ID from the context (set by AuthRequired middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to uint, assuming it is stored as a string in JWT
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Convert the string user ID to uint
	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

    var fcmToken models.FCMToken
    if err := config.DB.Where("user_id = ?", userIDUint).First(&fcmToken).Error; err == nil {
        fcmToken.Token = input.Token
        if err := config.DB.Save(&fcmToken).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update FCM token"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "FCM token updated successfully"})
        return
    }

    fcmToken = models.FCMToken{
		UserID: uint(userIDUint),
        Token:  input.Token,
    }

    if err := config.DB.Create(&fcmToken).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save FCM token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "FCM token saved successfully"})
}


// DeleteFCMToken akan menghapus token FCM yang terkait dengan pengguna yang sudah terautentikasi
func DeleteFCMToken(c *gin.Context) {
	// Get the user ID from the context (set by AuthRequired middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to uint, assuming it is stored as a string in JWT
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Convert the string user ID to uint
	userIDUint, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Cari dan hapus token FCM untuk user
	result := config.DB.Where("user_id = ?", userIDUint).Delete(&models.FCMToken{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No FCM token found for the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "FCM token deleted successfully"})
}
