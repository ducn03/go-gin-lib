package redis

import "os"

type Config struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig() Config {
	return Config{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // Hoặc lấy từ ENV nếu cần
		DB:       0,
	}
}
