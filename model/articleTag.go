package model

import (
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ArticleTagModel article_tag model
type ArticleTagModel struct {
	BaseModel
	TagID     uint64       `json:"tagId" gorm:"no null;column:tag_id"`
	Tag       TagModel     `json:"tag" gorm:"foreignKey:TagID"`
	ArticleID uint64       `json:"articleId" gorm:"no null;column:article_id"`
	Article   ArticleModel `json:"article" gorm:"foreignKey:ArticleID"`
}

// TableName article_tag table name
func (atm *ArticleTagModel) TableName() string {
	return "ws_article_tag"
}

func (atm *ArticleTagModel) Create() error {
	return DB.Self.Model(atm).Create(atm).Error
}

func CreateArticleTag(c *gin.Context, tagId uint64, articleId uint64) (error, *ArticleTagModel) {
	articleTag := &ArticleTagModel{
		TagID:     tagId,
		ArticleID: articleId,
	}
	err := articleTag.Create()
	if err != nil {
		logger.Logger.Error("articleTag insert failed", zap.String("error", err.Error()))
		return err, nil
	}
	return nil, articleTag
}

func QueryArticleTagListByArticleId(articleId uint64) (error, []ArticleTagModel) {
	cm := &ArticleTagModel{}
	list := []ArticleTagModel{}
	tx := DB.Self.Model(cm).Where("article_id=?", articleId).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("category list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryArticleTagListByTagId(tagId uint64) (error, []ArticleTagModel) {
	atm := &ArticleTagModel{}
	list := []ArticleTagModel{}
	tx := DB.Self.Model(atm).Preload("Tag").Preload("Article").Where("tag_id=?", tagId).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("articleTag list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}
