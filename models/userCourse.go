package models

import (
	"time"

	"gorm.io/gorm"
)

// UserCourse represents a course chosen by the user
type UserCourse struct {
    gorm.Model
    UserID     uint      `json:"user_id"`      // Relasi ke User
    CourseName string    `json:"course_name"`  // Nama mata kuliah
    Lecturer   string    `json:"lecturer"`     // Nama dosen
    StartTime  time.Time `json:"start_time"`   // Waktu mulai kuliah
    EndTime    time.Time `json:"end_time"`     // Waktu selesai kuliah
    Days       string    `json:"days"`         // Hari-hari kuliah (misalnya: "Senin, Rabu, Jumat")
    Classroom  string    `json:"classroom"`    // Ruang kelas
}
