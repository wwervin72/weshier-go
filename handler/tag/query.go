package tag

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// QueryTagList tag list query
func QueryTagList(c *gin.Context) {
	err, list := model.QueryTagList()
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

// QueryUserTagList tag list query
func QueryUserTagList(c *gin.Context) {
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, list := model.QueryUserTagList(user.ID)
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

// QueryTagInfo tag info query
func QueryTagInfo(c *gin.Context) {
	id := c.Param("tagId")
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
	err, tag := model.QueryTagByIDAndUser(uint64(intId), user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, tag)
	return
}

// QueryArticleStatisticByMonth article statistic
func QueryTagArticleStatistic(c *gin.Context) {
	err, statistic := model.QueryTagArticleStatistic()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}

// QueryArticleStatisticByMonth article statistic
func QueryUserTagArticleStatistic(c *gin.Context) {
	id := c.Param("userId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, statistic := model.QueryUserTagArticleStatistic(uint64(intId))
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}

// QueryTagPagination tag pagination
func QueryTagPagination(c *gin.Context) {
	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, tags, count := model.TagPagination(body)
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
			List:  tags,
			Total: count,
		},
	})
	return
}
