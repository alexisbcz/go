package env

import "os"

func GetVar(key, fallback string) string {
	if value, found := os.LookupEnv(key); found != false {
		return value
	}

	return fallback
}
