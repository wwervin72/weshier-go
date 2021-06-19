package model

import (
	"fmt"
	"weshierNext/pkg/auth"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/token"
	"weshierNext/pkg/validate"
	"weshierNext/util"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// UserModel user model
type UserModel struct {
	BaseModel
	UserName string `zh:"用户名" json:"userName" gorm:"not null;column:username;unique_index" binding:"required" validate:"min=3,max=10"`
	PassWord string `zh:"密码" json:"-" gorm:"not null;column:password;" binding:"required" validate:"min=5,max=50"`
	Email    string `zh:"邮箱" json:"email" gorm:"not null" binding:"required" validate:"email"`
	NickName string `zh:"昵称" json:"nickName" gorm:"column:nickname;" validate:"max=24"`
	Bio      string `zh:"简介" json:"bio"`
	Avatar   string `zh:"头像" json:"avatar"`
	URL      string `zh:"头像" json:"url"`
	Phone    uint64 `zh:"手机号" json:"phone"`
	Role     string `zh:"角色" json:"role"`
	Age      uint8  `zh:"年龄" json:"age"`
	Status   uint8  `json:"status"`
	Resume   uint8  `zh:"简历" json:"resume"`
	AuthID   uint64 `json:"authId" gorm:"column:auth_id;"`
}

// TableName specified table name
func (u *UserModel) TableName() string {
	return "ws_user"
}

// InsertAdminUser auto insert admin account into user table
func InsertAdminUser() (err error) {
	admin := &UserModel{
		UserName: viper.GetString("admin.userName"),
		PassWord: viper.GetString("admin.passWord"),
		Email:    viper.GetString("admin.email"),
		NickName: viper.GetString("admin.nickName"),
		Role:     viper.GetString("admin.role"),
	}
	if admin.NickName == "" {
		admin.NickName = util.RandomString(5, false)
	}
	err = admin.Validate()
	if err != nil {
		logger.Logger.Debug("admin user validate failed", zap.String("error", err.Error()))
		return
	}
	_, err = QueryUserByUsername(admin.UserName)
	if err == gorm.ErrRecordNotFound {
		err = admin.Create()
		if err != nil {
			logger.Logger.Panic("admin user insert failed", zap.String("error", err.Error()))
			return
		}
		return
	}
	if err != nil {
		logger.Logger.Panic("admin user query failed", zap.String("error", err.Error()))
	}
	return
}

// Create create a user
func (u *UserModel) Create() error {
	err := u.EncryptPwd()
	if err != nil {
		return err
	}
	return DB.Self.Model(u).Create(&u).Error
}

func (um *UserModel) AfterFind(tx *gorm.DB) (err error) {
	if um.Avatar == "" {
		um.Avatar = fmt.Sprintf("%s/asset/image/avatar.png", viper.GetString("publicPrefix"))
	}
	return
}

// Delete delete a user
func (u *UserModel) Delete() error {
	return DB.Self.Model(u).Delete(&u).Error
}

// DeleteByID delete a user by userID
func DeleteByID(id uint64) error {
	u := &UserModel{
		BaseModel: BaseModel{
			ID: id,
		},
	}
	return DB.Self.Model(u).Delete(&u).Error
}

// QueryUserByUsername query user by username
func QueryUserByUsername(username string) (UserModel, error) {
	u := UserModel{}
	data := DB.Self.Model(&u).Where("username=?", username).First(&u)
	return u, data.Error
}

// QueryUser 查询用户
func QueryUser(condition string, args ...interface{}) (*UserModel, error) {
	u := &UserModel{}
	data := DB.Self.Model(u).Where(condition, args...).First(u)
	return u, data.Error
}

// QueryUserByUserID 根据 ID 查询用户信息
func QueryUserByUserID(id uint64) (UserModel, error) {
	u := UserModel{}
	data := DB.Self.Model(&u).Where("id=?", id).First(&u)
	return u, data.Error
}

// Compare compare pwd whether same
func (u *UserModel) Compare(pwd string) error {
	err := auth.Compare(u.PassWord, pwd)
	return err
}

// EncryptPwd encry user password
func (u *UserModel) EncryptPwd() (err error) {
	u.PassWord, err = auth.Encrypt(u.PassWord)
	return err
}

// Validate Validate the field
func (u *UserModel) Validate() error {
	return validate.Validate(*u)
}

// Login 登录操作
func (u *UserModel) Login(c *gin.Context) (claims *token.JWTClaims, uid uuid.UUID, err error) {
	// sign the web token
	claims = &token.JWTClaims{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
		NickName: u.NickName,
		Bio:      u.Bio,
		Avatar:   u.Avatar,
		URL:      u.URL,
		Phone:    u.Phone,
		Role:     u.Role,
		Age:      u.Age,
		Status:   u.Status,
		Resume:   u.Resume,
		AuthID:   u.AuthID,
	}
	t, err := token.Sign(c, *claims, viper.GetString("jwt.secret"))
	if err != nil {
		logger.Logger.Debug("login generate token failed ", zap.String("error", err.Error()))
		return claims, uuid.Nil, errno.InternalServerError
	}
	// save token to redis
	redisConn := DB.RedisPool.Get()
	defer redisConn.Close()
	tokenKey, err := uuid.NewUUID()
	if err != nil {
		logger.Logger.Debug("login generate uuid failed", zap.String("error", err.Error()))
		return claims, uuid.Nil, errno.InternalServerError
	}
	// save token
	_, err = redisConn.Do("Set", tokenKey, t)
	if err != nil {
		logger.Logger.Debug("login save token failed", zap.String("error", err.Error()))
		return claims, uuid.Nil, errno.InternalServerError
	}
	// set expire time
	_, err = redisConn.Do("Expire", tokenKey, viper.GetInt64("jwt.maxage"))
	if err != nil {
		logger.Logger.Debug("login set token expire time failed", zap.String("error", err.Error()))
		return claims, uuid.Nil, errno.InternalServerError
	}
	return claims, tokenKey, nil
}

func GetLoginToken(tokenKey string) (bool, string, error) {
	redisConn := DB.RedisPool.Get()
	defer redisConn.Close()
	// 根据 key 找到 token
	t, err := redisConn.Do("GET", tokenKey)
	if err != nil {
		logger.Logger.Debug("query token failed", zap.String("error", err.Error()))
		return false, "", errno.InternalServerError
	}
	if t == nil {
		logger.Logger.Debug("query token failed", zap.String("error", "token is not found"))
		return false, "", errno.ErrNeedLogin
	}
	token, err := redis.String(t, nil)
	if err != nil {
		logger.Logger.Debug("convert token to string failed", zap.String("error", err.Error()))
		return false, "", errno.InternalServerError
	}
	// 以秒为单位，返回给定 key 的剩余生存时间
	timeRetain, err := redisConn.Do("TTL", tokenKey)
	if err != nil {
		logger.Logger.Debug("query token expire time failed", zap.String("error", err.Error()))
		return false, "", errno.InternalServerError
	}
	// 当 key 不存在时，返回 -2 。
	if timeRetain.(int64) == -2 {
		logger.Logger.Debug("TTL query return 2", zap.String("error", "token is not found"))
		return false, "", errno.ErrNeedLogin
	} else if timeRetain.(int64) == -1 {
		// 当 key 存在但没有设置剩余生存时间时，返回 -1 。
	} else if timeRetain.(int64) > 0 {
		// 否则，以秒为单位，返回 key 的剩余生存时间。
		// 刷新过期时间
		_, err = redisConn.Do("Expire", tokenKey, viper.GetInt64("jwt.maxage"))
		if err != nil {
			logger.Logger.Debug("refresh token expire time failed", zap.String("error", err.Error()))
			return false, "", errno.InternalServerError
		}
	}
	return true, token, nil
}

func ValidToken(tokenKey string) (bool, error) {
	// remove token in redis
	redisConn := DB.RedisPool.Get()
	defer redisConn.Close()
	// 以秒为单位，返回给定 key 的剩余生存时间
	timeRetain, err := redisConn.Do("TTL", tokenKey)
	// 当 key 不存在时，返回 -2 。
	if timeRetain.(int64) == -2 {
		logger.Logger.Debug("TTL query return 2", zap.String("error", "token is not found"))
		return false, errno.ErrNeedLogin
	} else if timeRetain.(int64) == -1 {
		// 当 key 存在但没有设置剩余生存时间时，返回 -1 。
	} else if timeRetain.(int64) > 0 {
		// 否则，以秒为单位，返回 key 的剩余生存时间。
		// 刷新过期时间
		_, err = redisConn.Do("Expire", tokenKey, viper.GetInt64("jwt.maxage"))
		if err != nil {
			logger.Logger.Debug("refresh token expire time failed", zap.String("error", err.Error()))
			return false, errno.InternalServerError
		}
	}
	return true, nil
}

// Logout 登出操作
func (u *UserModel) Logout(tokenKey string) (err error) {
	// remove token in redis
	redisConn := DB.RedisPool.Get()
	defer redisConn.Close()
	_, err = redisConn.Do("Del", tokenKey)
	if err != nil {
		logger.Logger.Debug("logout failed", zap.String("error", err.Error()))
		return errno.InternalServerError
	}
	return
}

func (u *UserModel) ChangePwd(oldPwd string) (err error) {
	aMap := make(map[string]interface{})
	aMap["password"] = u.PassWord
	aMap["id"] = u.ID
	var findUser UserModel
	tx := DB.Self.Model(u).Where("id=?", u.ID).Find(&findUser)
	if tx.Error != nil {
		logger.Logger.Debug("change password failed", zap.String("error", tx.Error.Error()))
		return errno.InternalServerError
	}
	if auth.Compare(findUser.PassWord, oldPwd) != nil {
		logger.Logger.Debug("change password failed", zap.String("error", "旧密码不正确"))
		return errno.ErrOldPasswordIncorrect
	}
	tx = DB.Self.Model(&findUser).Updates(aMap)
	if tx.Error != nil {
		logger.Logger.Debug("change password failed", zap.String("error", tx.Error.Error()))
		return errno.InternalServerError
	}
	return nil
}
