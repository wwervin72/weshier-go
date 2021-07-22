package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BaseModel base model define
type BaseModel struct {
	ID          uint64     `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	CreatedAt   *time.Time `json:"-"`
	CreatedTime string     `gorm:"-" json:"createdAt"`
	UpdatedAt   *time.Time `json:"-"`
	UpdatedTime string     `gorm:"-" json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
}

func (bm *BaseModel) AfterFind(tx *gorm.DB) error {
	return nil
}

// 文件上传目录名
const ThumbnailDir = "thumbnail"
const ImageDir = "upload"
const AnnexDir = "annex"
