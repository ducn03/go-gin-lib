package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func Connect() error {
	cfg := LoadConfig()

	// In ra màn hình để mình soi xem Docker nó nạp Env đúng chưa
	fmt.Printf("Đang thử kết nối Redis tại địa chỉ: [%s] (DB: %d)\n", cfg.Addr, cfg.DB)

	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Ping thử
	err := Client.Ping(Ctx).Err()
	if err != nil {
		// Bọc lỗi lại kèm thông tin chi tiết
		return fmt.Errorf("REDIS_ERROR: Không thể kết nối tới [%s]. Lỗi gốc: %w. ",
			cfg.Addr, err)
	}

	fmt.Printf("REDIS_OK: Kết nối thành công tới %s\n", cfg.Addr)
	return nil
}
