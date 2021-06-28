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

// LeaveComment
// @Summary 留言
// @Description 留言
// @Tags comment
// @Produce json
// @Param commentReqBody body request.CreateCommentReqStruct true "留言请求体内容"
// @Success 200 {object} handler.Response{data=model.CommentModel} "留言成功"
// @Router /api/comment [post]
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

// CommentArticle
// @Summary 评论文章
// @Description 评论文章
// @Tags comment
// @Produce json
// @Param commentArticle body request.CreateArticleCommentReqStruct true "评论文章参数"
// @Success 200 {object} handler.Response{data=model.CommentModel} "留言成功"
// @Router /api/comment/article [post]
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
