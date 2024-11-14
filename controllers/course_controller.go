package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
	"github.com/mfuadfakhruzaki/Jadwalin/models"
)

// CreateCourse handles creating a new course for a user
func CreateCourse(c *gin.Context) {
	var input struct {
		CourseName string    `json:"course_name" binding:"required"`
		Lecturer   string    `json:"lecturer" binding:"required"`
		StartTime  time.Time `json:"start_time" binding:"required"`
		EndTime    time.Time `json:"end_time" binding:"required"`
		Days       string    `json:"days" binding:"required"`
		Classroom  string    `json:"classroom" binding:"required"` // Added Classroom field
	}

	// Bind the input JSON to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
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

	// Create the UserCourse record
	userCourse := models.UserCourse{
		UserID:    uint(userIDUint),
		CourseName: input.CourseName,
		Lecturer:   input.Lecturer,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
		Days:       input.Days,
		Classroom:  input.Classroom, // Store classroom information
	}

	// Save course to the database
	if err := config.DB.Create(&userCourse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course created successfully", "course": userCourse})
}

// GetUserCourses fetches all courses for a user with caching in Redis
func GetUserCourses(c *gin.Context) {
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

    // Check Redis cache first
    cacheKey := fmt.Sprintf("user_courses:%d", userIDUint)
    cachedCourses, err := config.RedisClient.Get(context.Background(), cacheKey).Result()
    
    if err == nil && cachedCourses != "" {
        // Redis hit: return cached data
        var userCourses []models.UserCourse
        // Unmarshal the cached JSON into the userCourses slice
        if err := json.Unmarshal([]byte(cachedCourses), &userCourses); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached courses"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"courses": userCourses})
        return
    }

    // If not in cache, query the database
    var userCourses []models.UserCourse
    if err := config.DB.Where("user_id = ?", userIDUint).Find(&userCourses).Error; err != nil {
        // Handling case where database query fails
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
        return
    }

    // Cache the result in Redis with an expiration time (e.g., 10 minutes)
    coursesJSON, err := json.Marshal(userCourses)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal courses"})
        return
    }

    // Set the cache with an expiration time of 10 minutes
    if err := config.RedisClient.Set(context.Background(), cacheKey, coursesJSON, 10*time.Minute).Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
        return
    }

    // Return the courses for the user
    c.JSON(http.StatusOK, gin.H{"courses": userCourses})
}

