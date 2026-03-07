package server

import (
	testRoutes "go-gin-lib/internal/features/test/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.HealthHandler)

	apiPublic := r.Group("/api/public")
	testRoutes.RegisterRoutes(apiPublic)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) HealthHandler(c *gin.Context) {
	// Gọi hàm Health() từ service database mà ông đã khởi tạo lúc NewServer
	healthStatus := s.db.Health()

	// Nếu trong map trả về có lỗi, mình có thể trả về status 503 (Service Unavailable)
	// Hoặc đơn giản là trả về 200 kèm nội dung tình trạng để check cho dễ
	if status, ok := healthStatus["status"]; ok && status == "down" {
		c.JSON(http.StatusServiceUnavailable, healthStatus)
		return
	}

	c.JSON(http.StatusOK, healthStatus)
}
