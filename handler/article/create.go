package article

import (
	"fmt"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Create
// @Summary 创建文章
// @Description 创建文章
// @Tags article
// @Produce json
// @Param createArticle body request.CreateArticleReqStruct true "创建文章参数"
// @Success 200 {object} model.ArticleModel "创建文章"
// @Router /api/article [post]
func Create(c *gin.Context) {
	body := &request.CreateArticleReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	if body.Thumbnail == "" {
		body.Thumbnail = fmt.Sprintf("%s/asset/image/thumbnail.jpg", viper.GetString("publicPrefix"))
	}
	err, article := model.CreateArticle(c, body, user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, article)
	return
}
