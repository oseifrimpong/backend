package delivery

import (
	"backend/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatSvc struct {
	service service.ChatService
}

func NewChatController(service service.ChatService) *chatSvc {
	return &chatSvc{
		service: service,
	}
}

func (svc *chatSvc) ChatRequest(ctx *gin.Context) {
	// var req ChatRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.SecureJSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := svc.service.QueryChat(req.Question); err != nil {
	// 	ctx.SecureJSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	ctx.SecureJSON(http.StatusOK, "success")
}
