package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword mengenkripsi password menggunakan bcrypt
func EncryptPassword(password string) (string, error) {
    // Menggunakan bcrypt untuk mengenkripsi password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// VerifyPassword memverifikasi apakah password yang diberikan sesuai dengan hash yang tersimpan
func VerifyPassword(hashedPassword, password string) error {
    // Membandingkan password yang diberikan dengan hash
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return errors.New("invalid password")
    }
    return nil
}
