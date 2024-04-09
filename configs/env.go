package configs

import "os"

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		switch key {
		case "PORT":
			return "8000"
		case "VERSION":
			return "0.1.0"
		default:
			return ""
		}
	}
	return value
}
