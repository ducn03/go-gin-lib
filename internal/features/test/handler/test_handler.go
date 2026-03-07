package handler

import (
	responseUtils "go-gin-lib/internal/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
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
