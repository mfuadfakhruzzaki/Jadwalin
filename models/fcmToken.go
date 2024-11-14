// models/fcm_token.go
package models

import "gorm.io/gorm"

type FCMToken struct {
    ID     uint   `gorm:"primaryKey"`
    UserID uint   `gorm:"not null"`
    Token  string `gorm:"not null"`
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&FCMToken{})
}
