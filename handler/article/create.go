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

// Create create article
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
