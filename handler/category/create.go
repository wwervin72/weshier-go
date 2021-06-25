package category

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create
// @Summary 创建分类
// @Description 创建分类
// @Tags category
// @Produce json
// @Param createCate body request.CreateCategoryReqStruct true "创建分类参数"
// @Success 200 {object} model.CategoryModel "分类创建成功"
// @Router /api/cate [post]
func Create(c *gin.Context) {
	body := &request.CreateCategoryReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, cate := model.CreateCategory(c, body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, cate)
	return
}
