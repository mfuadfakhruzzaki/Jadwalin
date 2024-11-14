package models

import "gorm.io/gorm"

// User represents the structure of a user in the system
type User struct {
    gorm.Model
    FirstName  string       `json:"first_name" gorm:"not null"`
    LastName   string       `json:"last_name" gorm:"not null"`
    Username   string       `json:"username" gorm:"unique;not null"`
    Email      string       `json:"email" gorm:"unique;not null"`
    Password   string       `json:"-" gorm:"not null"`  // Store password as hash
    ProfilePic string       `json:"profile_pic"`
    Courses    []UserCourse `json:"courses"`
    FCMToken   []FCMToken   `json:"fcm_tokens"`  // Relasi ke FCMToken
}
