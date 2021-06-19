package model

import (
	"fmt"
	"weshierNext/pkg/errno"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// ThumbnailModel thumbnail model
type ThumbnailModel struct {
	BaseModel
	Name   string    `zh:"文件名" json:"name" gorm:"column:name"`
	Desc   string    `zh:"描述" json:"desc" gorm:"column:desc"`
	Size   int64     `zh:"大小" json:"size" gorm:"column:size;"`
	UserID uint64    `zh:"用户ID" json:"userId" gorm:"column:user_id;"`
	User   UserModel `zh:"用户" json:"user" validate:"-" gorm:"foreignKey:UserID"`
	Url    string    `zh:"图片地址" json:"url" validate:"-" gorm:"-"`
}

// TableName article table name
func (tm *ThumbnailModel) TableName() string {
	return "ws_thumbnail"
}

func (tm *ThumbnailModel) AfterFind(tx *gorm.DB) (err error) {
	tm.Url = fmt.Sprintf("%s/%s/%s", viper.GetString("publicPrefix"), ThumbnailDir, tm.Name)
	return
}

// Create create article
func (am *ThumbnailModel) Create() error {
	return DB.Self.Model(am).Create(&am).Error
}

// Delete delete article
func (am *ThumbnailModel) Delete() error {
	findTm := &ThumbnailModel{}
	tx := DB.Self.Where("id=? and user_id=?", am.ID, am.UserID).Find(findTm)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return errno.ErrNotFound
		}
		return errno.InternalServerError
	}
	tx = DB.Self.Delete(findTm)
	if tx.Error != nil {
		return errno.InternalServerError
	}
	return nil
}

// QueryThumbnailsByUser query thumbnail list of user
func QueryThumbnailsByUser(userId uint64) (error, *[]ThumbnailModel) {
	tm := &ThumbnailModel{}
	list := &[]ThumbnailModel{}
	data := DB.Self.Model(tm).Order("created_at desc").Where("user_id=?", userId).Find(list)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, list
}
