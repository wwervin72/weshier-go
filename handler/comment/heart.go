package comment

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

func HeartComment(c *gin.Context) {
	id := c.Param("commentId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	commentId := uint64(intId)
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	count, err := model.CountCommentHeartByIdAndUser(commentId, user.ID)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if count > 0 {
		handler.SendResponse(c, errno.CustomSuccessErrno("不能重复点赞"), nil)
		return
	}
	chm := &model.CommentHeartModel{
		UserID:    user.ID,
		CommentID: commentId,
	}
	err = chm.Create()
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("点赞成功"), nil)
	return
}

func CancelHeartComment(c *gin.Context) {
	id := c.Param("commentId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	commentId := uint64(intId)
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	count, err := model.CountCommentHeartByIdAndUser(commentId, user.ID)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if count == 0 {
		handler.SendResponse(c, errno.CustomSuccessErrno("你还未点赞该评论"), nil)
		return
	}
	chm := &model.CommentHeartModel{
		UserID:    user.ID,
		CommentID: commentId,
	}
	err = chm.Delete()
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("取消点赞成功"), nil)
	return
}
