package server

import (
	"backend/api/delivery"
	"backend/api/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	switch gin.Mode() {
	case gin.ReleaseMode:
		// logger = config.Logger()
	default:
		// logger = config.Logger()

		err := godotenv.Load()
		if err != nil {
			fmt.Println("error loading .env file")
		}
	}

	baseRoute := router.Group("/api")

	chatRouterGroup(baseRoute)
	uploadRouterGroup(baseRoute)

	return router
}

func chatRouterGroup(router *gin.RouterGroup) {

	svc := service.NewChatService()
	chatController := delivery.NewChatController(svc)

	chatGroup := router.Group("/chat")
	{
		chatGroup.POST("", chatController.ChatRequest)
		chatGroup.GET("")
		chatGroup.GET("/:id")
		chatGroup.PUT("/:id")
		chatGroup.DELETE("/:id")
	}
}

func uploadRouterGroup(router *gin.RouterGroup) {
	svc := service.NewUploadService()
	uploadController := delivery.NewUploadController(svc)

	uploadGroup := router.Group("/upload")
	{
		uploadGroup.POST("", uploadController.Upload)
	}
}
