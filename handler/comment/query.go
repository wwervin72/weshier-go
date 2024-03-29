package comment

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// CommentPagination
// @Summary 留言分页
// @Description 留言分页
// @Tags comment
// @Produce json
// @Param PaginationParam body request.PaginationReqStruct true "分页参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.CommentModel}} "留言分页查询成功"
// @Router /api/comment/page [get]
func CommentPagination(c *gin.Context) {
	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, comments, count := model.CommentPagination(body, 0)
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
			List:  comments,
			Total: count,
		},
	})
	return
}

// ArticleCommentPagination
// @Summary 文章评论分页
// @Description 文章评论分页
// @Tags comment
// @Produce json
// @Param articleId path integer true "文章id"
// @Param paginationParam body request.PaginationReqStruct true "分页参数"
// @Success 200 {object} handler.Response{data=response.PaginationDataStruct{list=[]model.CommentModel}} "文章留言查询成功"
// @Router /api/comment/article/{articleId}/page [get]
func ArticleCommentPagination(c *gin.Context) {
	id := c.Param("articleId")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	body := &request.PaginationReqStruct{}
	if err := c.ShouldBindQuery(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, comments, count := model.CommentPagination(body, uint64(articleId))
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
			List:  comments,
			Total: count,
		},
	})
	return
}
