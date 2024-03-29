package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Login
// @Summary 用户登录
// @Description 用户登录
// @Tags user
// @Param login body LoginReqStruct true "登录账号"
// @Produce json
// @Success 200 {object} handler.Response{data=LoginResStruct} "登录成功"
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	var u LoginReqStruct
	if err := c.ShouldBindJSON(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.QueryUser("username=? and auth_id=0", u.UserName)
	if err == gorm.ErrRecordNotFound {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	if err = user.Compare(u.PassWord); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	claims, uid, err := user.Login(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		Token:     uid.String(),
		JWTClaims: *claims,
	})
	return
}

// GithubLogin
// @Summary github 登录回调
// @Description github 登录回调
// @Tags user
// @Param code query string true "github 返回的 code"
// @Produce json
// @Success 200 {object} handler.Response{data=LoginResStruct} "github 登录成功"
// @Router /api/user/auth/github/callback [get]
func GithubLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 通过 code 在去请求 github 获取
	client := &http.Client{}
	// 创建新的 http request，自定义 Header
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s", viper.GetString("github.auth_url"),
		viper.GetString("github.client_id"), viper.GetString("github.client_secret"), code, viper.GetString("github.redirect_url")), nil)
	// 设置 accept-type
	request.Header.Add("accept", "application/json")
	if err != nil {
		logger.Logger.Error("github login failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		logger.Logger.Error("github login failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 注意关闭 body 防止内存泄漏
	defer resp.Body.Close()
	// 读取 res.body 中的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Error("github login failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	var data GithubAccessTokenRedirectStruct
	// 反序列化响应体内容，获取 access_token 内容
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Logger.Error("json Unmarshal failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 根据 access_token 去获取用户信息
	// 创建请求
	userReq, err := http.NewRequest("GET", fmt.Sprintf("%s?access_token=%s", viper.GetString("github.access_url"), data.AccessToken), nil)
	userReq.Header.Add("accept", "application/json")
	if err != nil {
		logger.Logger.Error("fetch github user info failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 发送请求
	userResp, err := client.Do(userReq)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	defer userResp.Body.Close()
	// 读取响应体里的数据
	userRespBody, err := ioutil.ReadAll(userResp.Body)
	if err != nil {
		logger.Logger.Error("red github user request body failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	var userRespData GithubUserInfoStruct
	err = json.Unmarshal(userRespBody, &userRespData)

	if err != nil {
		logger.Logger.Error("fetch github user request json unmarshal failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	userExiested, err := model.QueryUser("username=? and auth_id!=0", userRespData.UserName)
	userAuth := &model.UserAuth{}
	if err == gorm.ErrRecordNotFound {
		// 如果是第一次登录
		// 需要在本地数据库创建一个对应账号
		userExiested = &model.UserModel{
			UserName: userRespData.UserName,
			Email:    userRespData.Email,
			Bio:      userRespData.Bio,
			URL:      userRespData.URL,
			Avatar:   userRespData.Avatar,
			NickName: userRespData.Name,
			Role:     model.TOURIST,
		}
		userAuth.OpenID = userRespData.ID
		userAuth.LoginType = "github"
		userAuth.AccessToken = data.AccessToken
		tx := model.DB.Self.Begin()
		insertUserResult := tx.Create(userExiested)
		insertUserData, ok := insertUserResult.Value.(*model.UserModel)
		if insertUserResult.Error != nil || !ok {
			tx.Rollback()
			logger.Logger.Error("github login insert user failed", zap.String("error", err.Error()))
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		userAuth.UserID = insertUserData.ID
		insertUserAuthResult := tx.Create(userAuth)
		insertUserAuthData, ok := insertUserAuthResult.Value.(*model.UserAuth)
		if insertUserAuthResult.Error != nil || !ok {
			tx.Rollback()
			logger.Logger.Error("github login insert auth user failed", zap.String("error", err.Error()))
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		userExiested.AuthID = insertUserAuthData.ID
		// 更新 authID 字段
		err = tx.Save(&userExiested).Error
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("github login update auth id failed", zap.String("error", err.Error()))
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		// 所有操作都完成
		tx.Commit()
	} else if err != nil {
		logger.Logger.Error("github login query user failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	userAuth, err = model.QueryAuthByID(userExiested.AuthID)
	// 更新 access_token
	userAuth.AccessToken = data.AccessToken
	err = model.DB.Self.Save(&userAuth).Error
	if err != nil {
		logger.Logger.Error("github login update access_token failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	claims, uid, err := userExiested.Login(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		Token:     uid.String(),
		JWTClaims: *claims,
		UserAuth:  *userAuth,
		ID:        userExiested.ID,
	})
}
