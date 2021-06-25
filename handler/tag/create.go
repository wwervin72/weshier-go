package tag

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create
// @Summary 创建tag
// @Description 创建tag
// @Tags tag
// @Produce json
// @Param createTagParam body request.CreateTagReqStruct true "创建tag参数"
// @Success 200 {object} model.ArticleModel "tag创建成功"
// @Router /api/tag [post]
func Create(c *gin.Context) {
	body := &request.CreateTagReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, tag := model.CreateTag(body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, tag)
	return
}
