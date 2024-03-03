package util

import (
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var letters = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Generate ID generates random numbers for Id
func GenerateId() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	otp := rand.Intn(max-min+1) + min
	return int64(otp)
}

//ToSnakeCase is function to convert camelCase to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//SetLowerAndAddSpace is function to convert camelCase to snake_case
func SetLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

// IsDevelopmentEnv returns whether the app is running using production env
func IsDevelopmentEnv() bool {
	return os.Getenv("APP_ENV") == "development"
}

// RandSeq for
func RandSeq(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateRandomOTP generates random numbers for OTP
func GenerateRandomOTP() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	otp := rand.Intn(max-min+1) + min
	return int64(otp)
}
