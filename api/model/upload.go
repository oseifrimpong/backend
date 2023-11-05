package model

type Document struct {
	Base
	FileType string `json:"file_type" gorm:"column:file_type;type:varchar(30);"`
	FileName string `json:"file_name" gorm:"column:file_name;type:varchar(30);"`
	FileSize int64  `json:"file_size" gorm:"column:file_size;type:varchar(30);"`
	FileURL  string `json:"file_url" gorm:"column:file_url;type:varchar(30);"`
}
