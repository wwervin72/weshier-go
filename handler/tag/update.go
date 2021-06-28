package tag

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update
// @Summary 更新tag
// @Description 更新tag
// @Tags tag
// @Produce json
// @Param updateTagParam body request.UpdateTagReqStruct true "更新tag参数"
// @Success 200 {object} handler.Response{} "tag更新成功"
// @Router /api/tag [put]
func Update(c *gin.Context) {
	body := &request.UpdateTagReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err = model.UpdateTag(body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, errno.CustomSuccessErrno("tag更新成功"))
	return
}
