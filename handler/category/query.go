package category

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// QueryCategoryList
// @Summary 查询所有分类列表
// @Description 查询所有分类列表
// @Tags category
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.DropdownListResStruct{}} "查询分类列表成功"
// @Router /api/cate/list [get]
func QueryCategoryList(c *gin.Context) {
	err, list := model.QueryCategoryList()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	res := []response.DropdownListResStruct{}
	for _, item := range list {
		res = append(res, response.DropdownListResStruct{
			Label: item.Name,
			Value: item.ID,
		})
	}
	handler.SendResponse(c, nil, res)
	return
}

// QueryUserCategoryList
// @Summary 查询用户所有分类列表
// @Description 查询用户所有分类列表
// @Tags category
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.DropdownListResStruct{}} "查询用户分类列表成功"
// @Router /api/cate/user/list [get]
func QueryUserCategoryList(c *gin.Context) {
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	err, list := model.QueryUserCategoryList(user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	res := []response.DropdownListResStruct{}
	for _, item := range list {
		res = append(res, response.DropdownListResStruct{
			Label: item.Name,
			Value: item.ID,
		})
	}
	handler.SendResponse(c, nil, res)
	return
}

// QueryCatePagination
// @Summary 查询用户所有分类列表
// @Description 查询用户所有分类列表
// @Tags category
// @Produce json
// @Param PaginationParam body request.PaginationReqStruct true "分类分页查询参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.CategoryModel}} "查询用户分类分页成功"
// @Router /api/cate/page [get]
func QueryCatePagination(c *gin.Context) {
	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, catas, count := model.CataPagination(body)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &response.PaginationResStruct{
		Response: response.Response{
			Code:    errno.OK.Code,
			Message: errno.OK.Message,
		},
		Data: response.PaginationDataStruct{
			PaginationReqStruct: request.PaginationReqStruct{
				PageSize:   body.PageSize,
				PageNumber: body.PageNumber,
			},
			List:  catas,
			Total: count,
		},
	})
	return
}

// QueryCategoryInfo
// @Summary 查询分类详情
// @Description 查询分类详情
// @Tags category
// @Produce json
// @Param cataId path integer true "分类id"
// @Success 200 {object} handler.Response{data=model.CategoryModel} "分类详情查询成功"
// @Router /api/cate/detail/{cataId} [get]
func QueryCategoryInfo(c *gin.Context) {
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
	err, cata := model.QueryCategoryByIDAndUser(uint64(intId), user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, cata)
	return
}
