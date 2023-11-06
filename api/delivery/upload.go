package delivery

import (
	"backend/api/dto"
	"backend/api/service"
	"net/http"
	"os"

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

	file, _ := ctx.FormFile("file")
	// tmp := strings.Split(file.Filename, ".")
	// fileType := tmp[len(tmp)-1]

	var req dto.UploadRequest
	req.FileType = ctx.PostForm("file_type")

	if !validateFileType(req.FileType) {
		ctx.SecureJSON(http.StatusBadRequest, "invalid file type, only PDFs are accepted")
		return
	}

	err := ctx.SaveUploadedFile(file, os.Getenv("TMP_FILE_DIR")+file.Filename)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	req.FileURL = os.Getenv("TMP_FILE_DIR") + file.Filename

	//upload file to data folder
	if err := svc.service.Upload(req); err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "success")
}

func (svc *uploadSvc) Load(ctx *gin.Context) {

	var req dto.LoadRequest
	// req.FileURL = ctx.PostForm("file_url")

	if err := svc.service.Load(req); err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "success")
}

// only accept PDFs
func validateFileType(fileType string) bool {
	return fileType != "pdf" || fileType != "PDF"
}
