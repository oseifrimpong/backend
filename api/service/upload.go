package service

type UploadService interface {
}

type uploadSvc struct {
}

func NewUploadService() UploadService {
	return &uploadSvc{}
}

func (u *uploadSvc) SaveDoc() error {
	return nil
}
