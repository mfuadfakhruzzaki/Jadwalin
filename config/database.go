package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var DB *sqlx.DB

func InitDB() {
    // Ambil konfigurasi dari .env
    dbHost := GetEnv("DB_HOST", "localhost")
    dbPort := GetEnv("DB_PORT", "5432")
    dbUser := GetEnv("DB_USER", "user")
    dbPassword := GetEnv("DB_PASSWORD", "password")
    dbName := GetEnv("DB_NAME", "mydb")

    // Format string koneksi PostgreSQL
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
    var err error
    DB, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Database connected successfully")
}
