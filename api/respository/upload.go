package respository

import (
	"io"

	"gorm.io/gorm"
)

type UploadInt interface {
	SaveDoc() error
}

type upload struct {
	f  io.Reader
	db *gorm.DB
}

func NewUploadRepo(db *gorm.DB, f io.ReadCloser) UploadInt {
	return &upload{db: db, f: f}
}

func (u *upload) SaveDoc() error {
	return nil
}
