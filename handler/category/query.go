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

// QueryCategoryList category list query
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

// QueryCategoryList category list query
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

// QueryCataPagination cata pagination
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

// QueryTagInfo tag info query
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
	err, tag := model.QueryCategoryByIDAndUser(uint64(intId), user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, tag)
	return
}
