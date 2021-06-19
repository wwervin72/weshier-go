package model

import (
	"github.com/jinzhu/gorm"
)

// ThumbnailModel thumbnail model
type AnnexModel struct {
	BaseModel
	Name      string       `zh:"附件名" json:"name" gorm:"column:name"`
	Desc      string       `zh:"描述" json:"desc" gorm:"column:desc"`
	Size      int64        `zh:"大小" json:"size" gorm:"column:size;"`
	UserID    uint64       `zh:"用户ID" json:"userId" gorm:"column:user_id;"`
	User      UserModel    `zh:"用户" json:"user" validate:"-" gorm:"foreignKey:UserID"`
	ArticleID uint64       `zh:"文章ID" json:"articleId" gorm:"column:article_id;"`
	Article   ArticleModel `zh:"文章" json:"article" validate:"-" gorm:"foreignKey:ArticleID"`
	Url       string       `zh:"附件地址" json:"url" validate:"-" gorm:"-"`
}

// TableName article table name
func (am *AnnexModel) TableName() string {
	return "ws_annex"
}

// Create create annnex
func (am *AnnexModel) Create() error {
	return DB.Self.Model(am).Create(&am).Error
}

// QueryAnnexByUser query annex list of user
func QueryAnnexByUser(userId uint64) (error, *[]AnnexModel) {
	tm := &AnnexModel{}
	list := &[]AnnexModel{}
	data := DB.Self.Model(tm).Where("user_id=?", userId).Find(list)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, list
}

func QueryAnnexByIDAndUser(id uint64, userId uint64) (error, *AnnexModel) {
	tm := &AnnexModel{}
	data := DB.Self.Model(tm).Where("user_id = ? and id = ?", userId, id).Find(tm)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, tm
}
