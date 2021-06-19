package middleware

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// LoginRequired required user is logined
func LoginRequired(c *gin.Context) {
	tokenKey, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, errno.ErrNeedLogin, nil)
		c.Abort()
		return
	}
	tokenValid, tokenString, err := model.GetLoginToken(tokenKey)
	if !tokenValid || err != nil {
		handler.SendResponse(c, err, nil)
		c.Abort()
		return
	}
	// load the jwt secret from config
	secret := viper.GetString("jwt.secret")
	u, err := token.Parse(tokenString, secret)
	if err != nil {
		handler.SendResponse(c, err, nil)
		c.Abort()
		return
	}
	c.Set("user", u)
	c.Set("token", tokenKey)
	c.Next()
}

// IsAdmin whether user is admin role
func IsAdmin(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		handler.SendResponse(c, errno.InternalServerError, nil)
		c.Abort()
		return
	}
	user := u.(*token.JWTClaims)
	if user == nil {
		handler.SendResponse(c, errno.ErrNeedLogin, nil)
		c.Abort()
		return
	}
	if user.Role != model.ADMIN {
		handler.SendResponse(c, errno.ErrNoPermission, nil)
		c.Abort()
		return
	}
	c.Next()
}
