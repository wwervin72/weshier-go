package article

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update
// @Summary 更新文章
// @Description 更新文章
// @Tags article
// @Produce json
// @Param updateArticle body request.UpdateArticleReqStruct true "更新文章参数"
// @Success 200 {object} string "更新文章"
// @Router /api/article [put]
func Update(c *gin.Context) {
	body := &request.UpdateArticleReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err = model.UpdateArticle(c, body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
	return
}
