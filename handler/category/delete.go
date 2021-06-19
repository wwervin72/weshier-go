package category

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DeleteCatagory category delete
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
	return
}
