package config

import "time"

var (
    JWTSecretKey       = GetEnv("JWT_SECRET", "defaultsecretkey")
    JWTExpirationMinutes = 60  // default 60 menit
)

func GetJWTExpirationTime() time.Duration {
    return time.Duration(JWTExpirationMinutes) * time.Minute
}
