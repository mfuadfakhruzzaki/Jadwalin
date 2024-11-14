package utils

import (
	"errors"
	"log"

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
	// Debugging: Log hash yang disimpan dan password yang dimasukkan
	log.Printf("Length of stored hash: %d", len(hashedPassword))
    log.Printf("Length of input password: %d", len(password))


	// Verifikasi kecocokan antara password yang dimasukkan dengan hash yang tersimpan
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Log error jika password tidak cocok
		log.Printf("Error comparing password. Hash: %s, Password comparison failed.", hashedPassword)
		return errors.New("invalid password")
	}
	return nil
}
