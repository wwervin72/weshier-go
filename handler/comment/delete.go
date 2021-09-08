package comment

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DeleteComment
// @Summary 删除留言
// @Description 删除留言
// @Tags comment
// @Produce json
// @Param DelCommentReqStruct body request.DelCommentReqStruct true "删除留言请求体内容"
// @Success 200 {object} handler.Response{data=nil} "留言成功"
// @Router /api/comment [delete]
func DeleteComment(c *gin.Context) {
	body := &request.DelCommentReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	cm, err := model.QueryCommentById(body.CommentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			handler.SendResponse(c, errno.ErrNotFound, nil)
			return
		} else {
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
	} else {
		if cm.AuthorID != user.ID && user.Role != model.ADMIN {
			// 不是作者及管理员 则没有权限
			handler.SendResponse(c, errno.ErrNoPermission, nil)
			return
		} else {
			// 是作者 直接删除
			err = cm.Delete()
			if err != nil {
				handler.SendResponse(c, errno.InternalServerError, nil)
			}
		}
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("留言删除成功"), nil)
	return
}
