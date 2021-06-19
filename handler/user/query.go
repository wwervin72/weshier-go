package user

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// QueryUserInfo query userinfo by token
func QueryUserInfo(c *gin.Context) {
	err, u := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	auth := &model.UserAuth{}
	// 如果是第三方登录
	if u.AuthID != 0 {
		auth, err = model.QueryAuthByID(u.AuthID)
		if err == gorm.ErrRecordNotFound {
			handler.SendResponse(c, errno.ErrUserNotFound, nil)
			return
		} else if err != nil {
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
	}
	err, tokenKey := handler.GetTokenKeyFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		UserAuth:  *auth,
		JWTClaims: *u,
		ID:        u.ID,
		Token:     tokenKey,
	})
	return
}
