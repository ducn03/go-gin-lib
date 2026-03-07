package routes

import (
	testHandler "go-gin-lib/internal/features/test/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.RouterGroup) {
	testHandler := testHandler.NewTestHandler()
	register := route.Group("/test")
	{
		register.GET("", testHandler.TestResponse)
		register.GET("/redis", testHandler.TestRedis)
	}
}
