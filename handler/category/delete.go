package category

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DeleteCatagory
// @Summary 删除分类
// @Description 删除分类
// @Tags category
// @Produce json
// @Param cataId path integer true "分类id"
// @Success 200 {object} handler.Response{} "分类删除成功"
// @Router /api/cate/{cataId} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("cateId")
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
	err, cata := model.QueryCategoryByID(uint64(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			handler.SendResponse(c, err, nil)
			return
		}
		handler.SendResponse(c, err, nil)
		return
	}
	// 如果角色是管理员 则可以直接删除
	if user.Role == model.ADMIN {
		err = cata.Delete()
	} else {
		// 否则需要判断 user 是否匹配
		if cata.UserID == user.ID {
			err = cata.Delete()
		} else {
			handler.SendResponse(c, errno.ErrNotFound, nil)
			return
		}
	}
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, errno.CustomSuccessErrno("分类删除成功"))
	return
}
