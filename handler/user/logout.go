package user

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
)

// Logout
// @Summary 退出登录
// @Description 退出登录
// @Tags user
// @Produce json
// @Success 200 {object} string "退出登录"
// @Router /api/user/logout [get]
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
