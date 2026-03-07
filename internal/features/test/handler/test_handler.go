package handler

import (
	"go-gin-lib/internal/core/redis"
	responseUtils "go-gin-lib/internal/utils/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) TestRedis(c *gin.Context) {
	key := "user:profile:123"
	var source string

	// Để đo thời gian phản hồi
	start := time.Now()

	// 1. Kiểm tra trong Redis trước (Cache Hit?)
	result, err := redis.Get(key)

	if err == nil && result != "" {
		// --- TRƯỜNG HỢP 1: CACHE HIT (Có hàng trong Redis) ---
		source = "Lấy từ Redis (Fast Path)"
	} else {
		// --- TRƯỜNG HỢP 2: CACHE MISS (Chưa có hoặc lỗi) ---
		source = "Lấy từ Database (Slow Path)"

		// Giả lập Database đang bận xử lý mất 2 giây
		time.Sleep(2 * time.Second)
		valFromDB := "Dữ liệu Hồ sơ User #123 cực nặng từ MySQL"

		// Lấy được rồi thì tranh thủ cất vào Redis dùng cho lần sau
		// Để 30 giây cho dễ test lần 2
		_ = redis.Set(key, valFromDB, 30*time.Second)
		result = valFromDB
	}

	duration := time.Since(start)

	// 3. Trả về kết quả kèm theo thống kê để thấy sự khác biệt
	status := http.StatusOK
	responseUtils.Response(c, status, "Kết quả Test Cache", gin.H{
		"data":           result,
		"source":         source,
		"execution_time": duration.String(), // Chỗ này sẽ thấy 2s vs vài ms
		"instruction":    "Hãy thử F5 lần nữa trong vòng 30s để thấy tốc độ bàn thờ!",
	})
}

func (h *TestHandler) TestResponse(c *gin.Context) {
	scenario := c.Query("scenario")

	switch scenario {
	case "success":
		// Test 1: Thành công rực rỡ
		mockData := gin.H{
			"id":    100,
			"name":  "Gemini AI",
			"roles": []string{"admin", "developer"},
		}
		responseUtils.Success(c, mockData)

	case "business_error":
		// Test 2: Lỗi nghiệp vụ (VD: Sai OTP - 1001)
		// Mong muốn: HTTP Header 200, nhưng Body status 1001
		responseUtils.Error(c, 1001, "Mã xác thực của ông hết hạn rồi!")

	case "auth_error":
		// Test 3: Lỗi hệ thống chuẩn (401)
		// Mong muốn: HTTP Header 401, Body status 401
		responseUtils.Error(c, http.StatusUnauthorized, "Ông chưa đăng nhập nha")

	case "server_error":
		// Test 4: Lỗi 500
		responseUtils.Error(c, http.StatusInternalServerError, "Server đang bận đi nhậu rồi")

	default:
		// Test 5: Dùng hàm Response tổng quát (Trả về 201 Created)
		responseUtils.Response(c, http.StatusCreated, "Tạo mới thành công nè", gin.H{"id": 1})
	}
}
