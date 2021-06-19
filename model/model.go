package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BaseModel base model define
type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

func (bm *BaseModel) AfterFind(tx *gorm.DB) error {
	t, _ := time.Parse("2006-01-02 15:04:05", bm.CreatedAt.Format("2006-01-02 15:04:05"))
	bm.CreatedAt = &t
	return nil
	// (*bm.CreatedAt).Format("2006-01-02 15:04:05")
}

// 文件上传目录名
const ThumbnailDir = "thumbnail"
const ImageDir = "upload"
const AnnexDir = "annex"
