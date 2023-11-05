package delivery

import (
	"backend/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type uploadSvc struct {
	service service.UploadService
}

func NewUploadController(service service.UploadService) *uploadSvc {
	return &uploadSvc{
		service: service,
	}
}

func (svc *uploadSvc) Upload(ctx *gin.Context) {
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
