package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
	"github.com/mfuadfakhruzaki/Jadwalin/models"
	"github.com/mfuadfakhruzaki/Jadwalin/utils"
	"gorm.io/gorm"
)

// Register function to create a new user
func Register(c *gin.Context) {
	var input models.User

	// Bind the input JSON to User struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password using the EncryptPassword utility
	hashedPassword, err := utils.EncryptPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	input.Password = hashedPassword

	// Save user to the database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login function to authenticate the user and generate JWT
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	// Bind the input JSON to User struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find user by email or username
	if err := config.DB.Where("email = ? OR username = ?", input.Email, input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		}
		return
	}

	// Verify password using the VerifyPassword utility
	if err := utils.VerifyPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Convert user.ID (uint) to string before passing it to GenerateJWT
	userID := strconv.Itoa(int(user.ID))  // Convert uint to string

	// Generate JWT
	token, err := utils.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

// Logout handler
func Logout(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
        return
    }
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    // Store the token in the blacklist
    err := config.RedisClient.Set(context.Background(), "blacklist:"+tokenString, true, time.Duration(config.GetJWTExpirationTime())).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to blacklist token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}


// ResetPassword function to allow the user to reset their password
func ResetPassword(c *gin.Context) {
	var input struct {
		Email          string `json:"email"`
		OldPassword    string `json:"old_password"`
		NewPassword    string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	// Bind the input JSON to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate password confirmation
	if input.NewPassword != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Find the user by email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		}
		return
	}

	// Verify if the old password is correct using VerifyPassword utility
	if err := utils.VerifyPassword(user.Password, input.OldPassword); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old password"})
		return
	}

	// Hash the new password using EncryptPassword utility
	hashedPassword, err := utils.EncryptPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing new password"})
		return
	}

	// Update the password in the database
	user.Password = hashedPassword
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
