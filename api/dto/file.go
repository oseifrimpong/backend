package dto

type UploadRequest struct {
	FileType string `json:"fileType" binding:"required"`
	FileURL  string `json:"fileUrl" binding:"required"`
}

type LoadRequest struct {
	FileID string `json:"fileId" binding:"required"`
}
