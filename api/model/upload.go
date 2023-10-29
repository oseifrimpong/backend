package model

import (
	"io"

	"gorm.io/gorm"
)

type Document struct {
	Base
	FileType string `json:"file_type" gorm:"column:file_type;type:varchar(30);"`
	FileName string `json:"file_name" gorm:"column:file_name;type:varchar(30);"`
	FileSize int64  `json:"file_size" gorm:"column:file_size;type:varchar(30);"`
	FileURL  string `json:"file_url" gorm:"column:file_url;type:varchar(30);"`
}

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
