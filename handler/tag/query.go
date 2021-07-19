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

// QueryTagList
// @Summary tag列表查询
// @Description tag列表查询
// @Tags tag
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.DropdownListResStruct{}} "tag列表查询成功"
// @Router /api/tag/list [get]
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

// QueryUserTagList
// @Summary 用户tag列表查询
// @Description 用户tag列表查询
// @Tags tag
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.DropdownListResStruct{}} "用户tag列表查询成功"
// @Router /api/tag/user/list [get]
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

// QueryTagInfo
// @Summary tag详情查询
// @Description tag详情查询
// @Tags tag
// @Produce json
// @Success 200 {object} handler.Response{data=model.TagModel} "tag详情查询成功"
// @Router /api/tag/detail/{tagId} [get]
func QueryTagInfo(c *gin.Context) {
	id := c.Param("tagId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, tag := model.QueryTagByID(uint64(intId))
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, tag)
	return
}

// QueryUserTagInfo
// @Summary 用户tag详情查询
// @Description 用户tag详情查询
// @Tags tag
// @Produce json
// @Success 200 {object} handler.Response{data=model.TagModel} "用户tag详情查询成功"
// @Router /api/tag/detail/{tagId} [get]
func QueryUserTagInfo(c *gin.Context) {
	uId := c.Param("userId")
	userId, err := strconv.Atoi(uId)
	tId := c.Param("tagId")
	tagId, err := strconv.Atoi(tId)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, tag := model.QueryTagByIDAndUser(uint64(tagId), uint64(userId))
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, tag)
	return
}

// QueryTagArticleStatistic
// @Summary tag文章数按月统计查询
// @Description tag文章数按月统计查询
// @Tags tag
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.ArticleTagStatisticResStruct} "tag文章数按月统计查询成功"
// @Router /api/tag/statistic/article [get]
func QueryTagArticleStatistic(c *gin.Context) {
	err, statistic := model.QueryTagArticleStatistic()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}

// QueryUserTagArticleStatistic
// @Summary 用户tag文章数按月统计查询
// @Description 用户tag文章数按月统计查询
// @Tags tag
// @Produce json
// @Param userId path integer true "用户id"
// @Success 200 {object} handler.Response{data=[]response.ArticleTagStatisticResStruct} "用户tag按月统计查询成功"
// @Router /api/tag/statistic/article/{userId} [get]
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
	return
}

// QueryTagPagination
// @Summary tag分页查询
// @Description tag分页查询
// @Tags tag
// @Produce json
// @Param PaginationParam body request.PaginationReqStruct true "分页查询参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.TagModel}} "tag分页查询成功"
// @Router /api/tag/page [get]
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

// QueryTagArticlePagination
// @Summary tag下文章分页查询
// @Description tag下文章分页查询
// @Tags tag
// @Produce json
// @Param PaginationParam body request.PaginationReqStruct true "tag下文章分页查询参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.TagModel}} "tag文章分页查询成功"
// @Router /api/tag/{tagId}/article [get]
func QueryTagArticlePagination(c *gin.Context) {
	id := c.Param("tagId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, articles, count := model.TagArticlePagination(uint64(intId), body)
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
			List:  articles,
			Total: count,
		},
	})
	return
}
