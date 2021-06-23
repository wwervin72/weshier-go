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
// @Param Param path string true "查询文章详情参数"
// @Success 200 {object} model.ArticleModel "通过id查询文章详情"
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

// QueryArticle[1,2,3] article list pagination query
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

// QueryArticleStatisticByMonth article statistic
func QueryArticleStatisticByMonth(c *gin.Context) {
	err, statistic := model.QueryArticleStatisticByMonth()
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, statistic)
}

// QueryArticleStatisticByMonthAndUser article statistic
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
