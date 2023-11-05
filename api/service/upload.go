package service

type UploadService interface {
	SaveDoc() error
}

type uploadSvc struct {
}

func NewUploadService() UploadService {
	return &uploadSvc{}
}

func (u *uploadSvc) SaveDoc() error {
	return nil
}
