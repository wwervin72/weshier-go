package user

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/request"
	"weshierNext/pkg/auth"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// ChangePwd
// @Summary 修改密码
// @Description 修改密码
// @Tags user
// @Param changePwd body request.ChangeAccountPwdReqStruct true "修改密码"
// @Produce json
// @Success 200 {object} string "修改密码"
// @Router /api/user/changepwd [put]
func ChangePwd(c *gin.Context) {
	var data request.ChangeAccountPwdReqStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	err, u := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	// 如果是第三方登录账号 不允许修改
	if u.AuthID != 0 {
		handler.SendResponse(c, errno.ErrNoPermission, nil)
		return
	}
	encryptedNewPwd, err := auth.Encrypt(data.NewPwd)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	user := &model.UserModel{
		BaseModel: model.BaseModel{
			ID: u.ID,
		},
		PassWord: encryptedNewPwd,
	}
	err = user.ChangePwd(data.OldPwd)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, errno.CustomSuccessErrno("密码修改成功"), nil)
	return
}
