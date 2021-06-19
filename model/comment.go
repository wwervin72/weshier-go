package model

import (
	"weshierNext/model/request"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/validate"

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
	Comments    []CommentModel `json:"comments" validate:"-"`
	ReplyUserId uint64         `zh:"回复人" json:"replyUserId" gorm:"column:reply_user_id;"`
	ReplyUser   UserModel      `zh:"回复人" json:"replyUser" validate:"-" gorm:"foreignKey:ReplyUserId"`
}

// TableName comment table name
func (cm *CommentModel) TableName() string {
	return "ws_comment"
}

// Create create comment
func (cm *CommentModel) Create() error {
	return DB.Self.Model(cm).Create(&cm).Error
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
