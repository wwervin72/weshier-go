package file

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

func DeleteThumbnail(c *gin.Context) {
	id := c.Param("thumbnailId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	thumbnail := &model.ThumbnailModel{
		BaseModel: model.BaseModel{
			ID: uint64(intId),
		},
		UserID: user.ID,
	}
	err = thumbnail.Delete()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("缩略图删除成功"), nil)
}
