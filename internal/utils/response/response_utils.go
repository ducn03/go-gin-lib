package response

import (
	"go-gin-lib/internal/core/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success - Ngắn nhất, dùng cho 90% trường hợp thành công (200 OK)
func Success(c *gin.Context, data interface{}) {
	res := dto.SuccessResponse(data)
	c.JSON(200, res)
}

// Error - Ngắn gọn cho các trường hợp lỗi, chỉ cần status và message
func Error(c *gin.Context, status int, message string) {
	httpStatus := getValidHttpStatus(status)

	res := dto.ErrorResponse(status, message)
	c.JSON(httpStatus, res)
}

// Response - Hàm tổng quát nhất, dùng khi cần tùy biến hoàn toàn (status, message, data)
func Response(c *gin.Context, status int, message string, data interface{}) {
	// Tận dụng constructor
	// Hoặc khởi tạo trực tiếp qua Success/Error nếu muốn phân luồng
	if status == 200 {
		res := dto.SuccessResponse(data)
		c.JSON(status, res)
		return
	}
	Error(c, status, message)
}

func getValidHttpStatus(status int) int {
	switch status {
	case http.StatusBadRequest, // 400
		http.StatusUnauthorized,        // 401
		http.StatusForbidden,           // 403
		http.StatusNotFound,            // 404
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable:  // 503
		return status
	default:
		// Mọi mã lỗi nghiệp vụ khác hoặc mã không nằm trong list trên
		// đều trả về HTTP 200 để FE tự xử lý dựa trên meta.status trong body
		return http.StatusOK
	}
}
