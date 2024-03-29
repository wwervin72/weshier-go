package article

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// QueryArticleDetailByID
// @Summary 查询文章详情
// @Description 通过id查询文章详情
// @Tags article
// @Produce json
// @Param articleId path integer true "文章id"
// @Success 200 {object} handler.Response{data=model.ArticleModel} "文章详情查询成功"
// @Router /api/article/detail/{articleId} [get]
func QueryArticleDetailByID(c *gin.Context) {
	id := c.Param("articleId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, article := model.QueryArticleByID(uint64(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			handler.SendResponse(c, errno.ErrNotFound, nil)
			return
		}
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, article)
	return
}

// QueryArticleList
// @Summary 文章分页查询
// @Description 文章分页查询
// @Tags article
// @Produce json
// @Param PaginationParam body request.PaginationReqStruct true "分页查询参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.ArticleModel}} "文章分页查询成功"
// @Router /api/article/page [get]
func QueryArticleList(c *gin.Context) {
	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, articles, count := model.ArticlePagination(body)
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

// QueryArticleStatisticByMonth
// @Summary 文章按月统计
// @Description 文章按月统计
// @Tags article
// @Produce json
// @Success 200 {object} handler.Response{data=[]response.ArticleMonthStatisticResStruct} "文章按月统计查询成功"
// @Router /api/article/statistic/month [get]
func QueryArticleStatisticByMonth(c *gin.Context) {
	err, statistic := model.QueryArticleStatisticByMonth()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}

// QueryArticleStatisticByMonthAndUser
// @Summary 用户文章按月统计
// @Description 用户文章按月统计
// @Tags article
// @Produce json
// @Param userId path integer true "用户id"
// @Success 200 {object} handler.Response{data=[]response.ArticleMonthStatisticResStruct} "用户文章按月统计查询成功"
// @Router /api/article/statistic/month/{userId} [get]
func QueryArticleStatisticByMonthAndUser(c *gin.Context) {
	id := c.Param("userId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	err, statistic := model.QueryArticleStatisticByMonthAndUser(uint64(intId))
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}
