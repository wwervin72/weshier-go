package model

import (
	"weshierNext/model/request"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/validate"
	"weshierNext/util"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// CommentModel article model
type CommentModel struct {
	BaseModel
	Content     string         `zh:"内容" json:"content" gorm:"not null;column:content" validate:"required,min=1"`
	AuthorID    uint64         `zh:"作者" json:"authorId" gorm:"not null;column:author_id" validate:"required"`
	Author      UserModel      `zh:"作者" json:"author" validate:"-" gorm:"foreignKey:AuthorID"`
	ArticleID   uint64         `zh:"文章" json:"articleId" gorm:"column:article_id;"`
	Article     ArticleModel   `json:"article" validate:"-" gorm:"foreignKey:ArticleID"`
	CommentID   uint64         `zh:"文章" json:"commentId" gorm:"column:comment_id;"`
	Comments    []CommentModel `json:"comments" validate:"-" gorm:"-"`
	ReplyUserId uint64         `zh:"回复人" json:"replyUserId" gorm:"column:reply_user_id;"`
	ReplyUser   UserModel      `zh:"回复人" json:"replyUser" validate:"-" gorm:"foreignKey:ReplyUserId"`
	HeartCount  uint16         `zh:"点赞数" json:"heartCount" gorm:"-"`
	HeartUserID []uint64       `zh:"点赞用户" json:"heartUserID" gorm:"-"`
}

// TableName comment table name
func (cm *CommentModel) TableName() string {
	return "ws_comment"
}

// Create create comment
func (cm *CommentModel) Create() error {
	return DB.Self.Model(cm).Create(&cm).Error
}

func (am *CommentModel) AfterFind(tx *gorm.DB) (err error) {
	am.CreatedTime = util.TimeFormat(*am.CreatedAt)
	am.UpdatedTime = util.TimeFormat(*am.UpdatedAt)
	userIds, err := QueryCommentHeartUserById(am.ID)
	if err != nil {
		return err
	}
	am.HeartCount = uint16(len(userIds))
	am.HeartUserID = userIds
	return nil
}

// QueryById query comment detail by id
func QueryCommentById(commentId uint64) (*CommentModel, error) {
	cm := &CommentModel{}
	tx := DB.Self.Model(cm).Where("id = ?", commentId).Find(cm)
	return cm, tx.Error
}

// Delete delete comment
func (cm *CommentModel) Delete() error {
	return DB.Self.Model(cm).Delete(&cm).Error
}

// Validate Validate the field
func (cm *CommentModel) Validate() error {
	return validate.Validate(*cm)
}

func CommentPagination(data *request.PaginationReqStruct, articleId uint64) (error, []CommentModel, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]CommentModel, pagesize)
	cm := &CommentModel{}
	var count uint
	tx := DB.Self.Model(cm).Where("article_id=? and comment_id=?", articleId, 0).Count(&count).Order("created_at desc").Limit(pagesize).Offset(offset).Preload("Author").Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("comment pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.InternalServerError, nil, 0
	}
	// 查询 reply
	for index, item := range list {
		err, comments, _ := QueryCommentReplyPagination(articleId, item.ID, data)
		if err != nil {
			logger.Logger.Error("comment pagination query failed", zap.String("error", tx.Error.Error()))
			return errno.InternalServerError, nil, 0
		}
		list[index].Comments = comments
	}
	return nil, list, count
}

func QueryArticleCommentCount(articleId uint64) (count uint16, err error) {
	cm := &CommentModel{}
	tx := DB.Self.Model(cm).Where("article_id = ?", articleId).Count(&count)
	return count, tx.Error
}

func QueryCommentReplyPagination(articleId, commentId uint64, data *request.PaginationReqStruct) (error, []CommentModel, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]CommentModel, pagesize)
	cm := &CommentModel{}
	var count uint
	tx := DB.Self.Model(cm).Where("comment_id=? and article_id=?", commentId, articleId).Count(&count).Order("created_at desc").Limit(pagesize).Offset(offset).Preload("Author").Preload("ReplyUser").Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("comment pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.InternalServerError, nil, 0
	}
	return nil, list, count
}
