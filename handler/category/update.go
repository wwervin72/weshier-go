package category

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update
// @Summary 更新分类
// @Description 更新分类
// @Tags category
// @Produce json
// @Param updateCate body request.UpdateCategoryReqStruct true "更新分类参数"
// @Success 200 {object} handler.Response{} "分类更新成功"
// @Router /api/cate [put]
func Update(c *gin.Context) {
	body := &request.UpdateCategoryReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err = model.UpdateCategory(body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, errno.CustomSuccessErrno("分类更新成功"))
	return
}
