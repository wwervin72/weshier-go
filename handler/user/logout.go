package user

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	tokenKey, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	tokenValid, _, err := model.GetLoginToken(tokenKey)
	if !tokenValid || err != nil {
		handler.SendResponse(c, err, nil)
		c.Abort()
		return
	}
	user := &model.UserModel{}
	err = user.Logout(tokenKey)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("登出成功"), nil)
	return
}
