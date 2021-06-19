package comment

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LeaveComment leave comment fot platform
func LeaveComment(c *gin.Context) {
	body := &request.CreateCommentReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	comment := &model.CommentModel{
		Content:  body.Content,
		AuthorID: user.ID,
	}
	if body.CommentID != 0 {
		comment.CommentID = body.CommentID
	}
	err = comment.Create()
	if err != nil {
		logger.Logger.Error("comment create failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, nil, comment)
	return
}

// CommentArticle leave comment fot article
func CommentArticle(c *gin.Context) {
	body := &request.CreateArticleCommentReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	comment := &model.CommentModel{
		Content:   body.Content,
		AuthorID:  user.ID,
		ArticleID: body.ArticleID,
	}
	if body.CommentID != 0 {
		comment.CommentID = body.CommentID
	}
	err = comment.Create()
	if err != nil {
		logger.Logger.Error("comment create failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, nil, comment)
	return
}
