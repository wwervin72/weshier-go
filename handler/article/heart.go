package article

import (
	"strconv"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// HeartArticle
// @Summary 点赞文章
// @Description 点赞文章
// @Tags article
// @Produce json
// @Success 200 {object} handler.Response{} "文章点赞成功"
// @Router /api/article/heart/{articleId} [get]
func HeartArticle(c *gin.Context) {
	id := c.Param("articleId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	articleId := uint64(intId)
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	count, err := model.CountArticleHeartByIdAndUser(articleId, user.ID)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if count > 0 {
		handler.SendResponse(c, errno.CustomSuccessErrno("不能重复点赞"), nil)
		return
	}
	chm := &model.ArticleHeartModel{
		UserID:    user.ID,
		ArticleID: articleId,
	}
	err = chm.Create()
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("点赞成功"), nil)
	return
}

// CancelHeartArticle
// @Summary 取消点赞文章
// @Description 取消点赞文章
// @Tags article
// @Produce json
// @Success 200 {object} handler.Response{} "取消点赞文章成功"
// @Router /api/article/heart/{articleId} [delete]
func CancelHeartArticle(c *gin.Context) {
	id := c.Param("articleId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	articleId := uint64(intId)
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	count, err := model.CountArticleHeartByIdAndUser(articleId, user.ID)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if count == 0 {
		handler.SendResponse(c, errno.CustomSuccessErrno("你还未点赞该评论"), nil)
		return
	}
	chm := &model.ArticleHeartModel{
		UserID:    user.ID,
		ArticleID: articleId,
	}
	err = chm.Delete()
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("取消点赞成功"), nil)
	return
}
